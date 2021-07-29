package flusher_test

import (
	"fmt"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ozoncp/ocp-survey-api/internal/flusher"
	"github.com/ozoncp/ocp-survey-api/internal/mocks"
	"github.com/ozoncp/ocp-survey-api/internal/models"
)

var _ = Describe("Flusher", func() {

	var (
		mockCtrl *gomock.Controller
		mockRepo *mocks.MockRepo
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(mockCtrl)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("Create new Flusher", func() {

		Context("with invalid arguments", func() {

			When("no Repo is specified", func() {
				It("returns an error", func() {
					f, err := flusher.New(2, nil)
					Expect(f).Should(BeNil())
					Expect(err).Should(HaveOccurred())
				})
			})

			When("invalid chunk size is specified", func() {
				It("returns an error", func() {
					f, err := flusher.New(0, mockRepo)
					Expect(f).Should(BeNil())
					Expect(err).Should(HaveOccurred())
				})
			})
		})

		Context("with correct arguments", func() {
			It("returns a Flusher instance", func() {
				f, err := flusher.New(2, mockRepo)
				Expect(f).ShouldNot(BeNil())
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
	})

	Describe("Flush", func() {

		var (
			data []models.Survey
		)

		BeforeEach(func() {
			data = []models.Survey{
				{Id: 0}, {Id: 1}, {Id: 2}, {Id: 3},
				{Id: 4}, {Id: 5}, {Id: 6}, {Id: 7},
			}
		})

		Context("successful save to Repo", func() {

			When("data is split into multiple chunks", func() {
				It("should flush all items", func() {
					f, _ := flusher.New(4, mockRepo)

					gomock.InOrder(
						mockRepo.EXPECT().AddSurvey(data[:4]),
						mockRepo.EXPECT().AddSurvey(data[4:]),
					)

					r, err := f.Flush(data)
					Expect(r).Should(BeNil())
					Expect(err).ShouldNot(HaveOccurred())
				})
			})

			When("there is only one chunk", func() {
				It("should flush all items", func() {
					f, _ := flusher.New(4, mockRepo)

					mockRepo.EXPECT().AddSurvey(data[:3])

					r, err := f.Flush(data[:3])
					Expect(r).Should(BeNil())
					Expect(err).ShouldNot(HaveOccurred())
				})
			})
		})

		Context("save to Repo failed", func() {

			When("partially saved", func() {
				It("should return remaining items", func() {
					f, _ := flusher.New(2, mockRepo)

					gomock.InOrder(
						mockRepo.EXPECT().AddSurvey(data[:2]),
						mockRepo.EXPECT().AddSurvey(data[2:4]),
						mockRepo.EXPECT().AddSurvey(data[4:6]).Return(fmt.Errorf("repo error")),
					)

					r, err := f.Flush(data)
					Expect(r).Should(BeEquivalentTo(data[4:]))
					Expect(err).Should(HaveOccurred())
				})
			})

			When("no data saved", func() {
				It("should return all items", func() {
					f, _ := flusher.New(4, mockRepo)

					mockRepo.EXPECT().AddSurvey(data[:4]).Return(fmt.Errorf("repo error"))

					r, err := f.Flush(data)
					Expect(r).Should(BeEquivalentTo(data))
					Expect(err).Should(HaveOccurred())
				})
			})
		})
	})
})

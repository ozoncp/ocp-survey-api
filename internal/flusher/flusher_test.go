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
		Context("with correct arguments", func() {
			It("returns a Flusher instance", func() {
				f := flusher.New(2, mockRepo)
				Expect(f).ShouldNot(BeNil())
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
					f := flusher.New(4, mockRepo)

					gomock.InOrder(
						mockRepo.EXPECT().AddSurvey(data[:4]),
						mockRepo.EXPECT().AddSurvey(data[4:]),
					)

					r := f.Flush(data)
					Expect(r).Should(BeNil())
				})
			})

			When("there is only one chunk", func() {
				It("should flush all items", func() {
					f := flusher.New(4, mockRepo)

					mockRepo.EXPECT().AddSurvey(data[:3])

					r := f.Flush(data[:3])
					Expect(r).Should(BeNil())
				})
			})
		})

		Context("save to Repo failed", func() {

			When("partially saved", func() {
				It("should return remaining items", func() {
					f := flusher.New(2, mockRepo)

					gomock.InOrder(
						mockRepo.EXPECT().AddSurvey(data[:2]),
						mockRepo.EXPECT().AddSurvey(data[2:4]),
						mockRepo.EXPECT().AddSurvey(data[4:6]).Return(fmt.Errorf("repo error")),
					)

					r := f.Flush(data)
					Expect(r).Should(BeEquivalentTo(data[4:]))
				})
			})

			When("no data saved", func() {
				It("should return all items", func() {
					f := flusher.New(4, mockRepo)

					mockRepo.EXPECT().AddSurvey(data[:4]).Return(fmt.Errorf("repo error"))

					r := f.Flush(data)
					Expect(r).Should(BeEquivalentTo(data))
				})
			})
		})

		Context("invalid input", func() {
			When("empty slice passed", func() {
				It("should return empty slice", func() {
					f := flusher.New(2, mockRepo)

					r := f.Flush(data[0:0])
					Expect(r).Should(BeEmpty())
					r = f.Flush(nil)
					Expect(r).Should(BeEmpty())
				})
			})
		})
	})
})

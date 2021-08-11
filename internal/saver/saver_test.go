package saver_test

import (
	"context"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ozoncp/ocp-survey-api/internal/mocks"
	"github.com/ozoncp/ocp-survey-api/internal/models"
	"github.com/ozoncp/ocp-survey-api/internal/saver"
)

var _ = Describe("Saver", func() {

	var (
		mockCtrl    *gomock.Controller
		mockFlusher *mocks.MockFlusher
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockFlusher = mocks.NewMockFlusher(mockCtrl)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("Create new Saver", func() {
		Context("with correct arguments", func() {
			It("returns a Saver instance", func() {
				svr := saver.New(context.Background(), 3, mockFlusher, 1*time.Second)
				Expect(svr).ShouldNot(BeNil())
				svr.Close()
			})
		})
	})

	Describe("Save", func() {

		var (
			svr      saver.Saver
			data     []models.Survey
			capacity uint          = 3
			timeout  time.Duration = 1 * time.Second
		)

		BeforeEach(func() {
			svr = saver.New(context.Background(), capacity, mockFlusher, timeout)

			data = []models.Survey{
				{Id: 0}, {Id: 1}, {Id: 2}, {Id: 3},
				{Id: 4}, {Id: 5}, {Id: 6}, {Id: 7},
			}
		})

		When("items exceed capacity", func() {
			It("should flush all items in batches", func() {
				gomock.InOrder(
					mockFlusher.EXPECT().Flush(data[0:capacity]),
					mockFlusher.EXPECT().Flush(data[capacity:capacity*2]),
					mockFlusher.EXPECT().Flush(data[capacity*2:]),
				)

				for _, survey := range data {
					svr.Save(survey)
				}
				svr.Close()
			})
		})

		When("less items than capacity", func() {
			It("should flush all items on close", func() {
				mockFlusher.EXPECT().Flush(data[:capacity-1])

				for _, survey := range data[:capacity-1] {
					svr.Save(survey)
				}
				svr.Close()
			})
		})
	})
})

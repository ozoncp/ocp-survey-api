package api_test

import (
	"context"
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opentracing/opentracing-go/mocktracer"

	"github.com/ozoncp/ocp-survey-api/internal/api"
	"github.com/ozoncp/ocp-survey-api/internal/mocks"
	"github.com/ozoncp/ocp-survey-api/internal/producer"
	"github.com/ozoncp/ocp-survey-api/internal/repo"
	desc "github.com/ozoncp/ocp-survey-api/pkg/ocp-survey-api"
)

var _ = Describe("Survey Service API", func() {

	var (
		rep repo.Repo
		srv desc.OcpSurveyApiServer

		ctx    context.Context
		cancel context.CancelFunc
		db     *sqlx.DB

		sqlm   sqlmock.Sqlmock
		ctrl   *gomock.Controller
		prod   *mocks.MockProducer
		metr   *mocks.MockMetrics
		tracer *mocktracer.MockTracer
		topic  = "survey_events"
	)

	BeforeEach(func() {
		ctx, cancel = context.WithCancel(context.Background())
		sqldb, sqlmck, err := sqlmock.New()
		if err != nil {
			panic(err)
		}
		sqlm = sqlmck
		db = sqlx.NewDb(sqldb, "")
		rep = repo.NewSurveyRepo(db)

		ctrl = gomock.NewController(GinkgoT())
		prod = mocks.NewMockProducer(ctrl)
		metr = mocks.NewMockMetrics(ctrl)
		tracer = mocktracer.New()

		srv = api.NewOcpSurveyApi(rep, prod, metr, tracer, 2)
	})

	AfterEach(func() {
		db.Close()
		ctrl.Finish()
		cancel()
	})

	Describe("CreateSurveyV1", func() {

		It("should store item to repo", func() {
			data := &desc.CreateSurveyV1Request{
				UserId: 1,
				Link:   "http://api.test/survey/1",
			}

			sqlm.ExpectQuery("INSERT INTO surveys").
				WithArgs(data.UserId, data.Link).
				WillReturnRows(sqlm.NewRows([]string{"id"}).AddRow(1))

			ev := producer.PrepareEvent(producer.Create, 1)
			prod.EXPECT().Send(topic, ev)
			metr.EXPECT().IncCreate()

			resp, err := srv.CreateSurveyV1(ctx, data)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp).ShouldNot(BeNil())
			Expect(resp.GetSurveyId()).Should(BeEquivalentTo(uint64(1)))

			Expect(sqlm.ExpectationsWereMet()).ShouldNot(HaveOccurred())
		})
	})

	Describe("MultiCreateSurveyV1", func() {

		When("multiple items passed", func() {
			It("should store items to repo", func() {
				data := []*desc.CreateSurveyV1Request{
					{UserId: 1, Link: "http://api.test/survey/1"},
					{UserId: 2, Link: "http://api.test/survey/2"},
					{UserId: 3, Link: "http://api.test/survey/3"},
				}
				req := desc.MultiCreateSurveyV1Request{
					Surveys: data,
				}

				sqlm.ExpectBegin()
				prep := sqlm.ExpectPrepare("INSERT INTO surveys")
				prep.ExpectQuery().
					WithArgs(data[0].UserId, data[0].Link).
					WillReturnRows(sqlm.NewRows([]string{"id"}).AddRow(1))
				prep.ExpectQuery().
					WithArgs(data[1].UserId, data[1].Link).
					WillReturnRows(sqlm.NewRows([]string{"id"}).AddRow(2))
				sqlm.ExpectCommit()
				sqlm.ExpectQuery("INSERT INTO surveys").
					WithArgs(data[2].UserId, data[2].Link).
					WillReturnRows(sqlm.NewRows([]string{"id"}).AddRow(3))

				ev := producer.PrepareEvent(producer.Create, 1)
				prod.EXPECT().Send(topic, ev)
				metr.EXPECT().IncCreate()
				ev = producer.PrepareEvent(producer.Create, 2)
				prod.EXPECT().Send(topic, ev)
				metr.EXPECT().IncCreate()
				ev = producer.PrepareEvent(producer.Create, 3)
				prod.EXPECT().Send(topic, ev)
				metr.EXPECT().IncCreate()

				resp, err := srv.MultiCreateSurveyV1(ctx, &req)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(resp).ShouldNot(BeNil())
				Expect(resp.GetSurveyIds()).Should(BeEquivalentTo([]uint64{1, 2, 3}))

				Expect(sqlm.ExpectationsWereMet()).ShouldNot(HaveOccurred())
			})
		})

		When("store to repo partially failed", func() {
			It("should return IDs of stored items", func() {
				data := []*desc.CreateSurveyV1Request{
					{UserId: 1, Link: "http://api.test/survey/1"},
					{UserId: 2, Link: "http://api.test/survey/2"},
					{UserId: 3, Link: "http://api.test/survey/3"},
				}
				req := desc.MultiCreateSurveyV1Request{
					Surveys: data,
				}

				sqlm.ExpectBegin()
				prep := sqlm.ExpectPrepare("INSERT INTO surveys")
				prep.ExpectQuery().
					WithArgs(data[0].UserId, data[0].Link).
					WillReturnRows(sqlm.NewRows([]string{"id"}).AddRow(1))
				prep.ExpectQuery().
					WithArgs(data[1].UserId, data[1].Link).
					WillReturnRows(sqlm.NewRows([]string{"id"}).AddRow(2))
				sqlm.ExpectCommit()
				sqlm.ExpectQuery("INSERT INTO surveys").
					WithArgs(data[2].UserId, data[2].Link).
					WillReturnError(sql.ErrNoRows)

				ev := producer.PrepareEvent(producer.Create, 1)
				prod.EXPECT().Send(topic, ev)
				metr.EXPECT().IncCreate()
				ev = producer.PrepareEvent(producer.Create, 2)
				prod.EXPECT().Send(topic, ev)
				metr.EXPECT().IncCreate()

				resp, err := srv.MultiCreateSurveyV1(ctx, &req)
				Expect(err).Should(HaveOccurred())
				Expect(resp).ShouldNot(BeNil())
				Expect(resp.GetSurveyIds()).Should(BeEquivalentTo([]uint64{1, 2}))

				Expect(sqlm.ExpectationsWereMet()).ShouldNot(HaveOccurred())
			})
		})

		When("single item passed", func() {
			It("should store item to repo", func() {
				data := &desc.CreateSurveyV1Request{
					UserId: 1,
					Link:   "http://api.test/survey/1",
				}
				req := desc.MultiCreateSurveyV1Request{
					Surveys: []*desc.CreateSurveyV1Request{data},
				}

				sqlm.ExpectQuery("INSERT INTO surveys").
					WithArgs(data.UserId, data.Link).
					WillReturnRows(sqlm.NewRows([]string{"id"}).AddRow(1))

				ev := producer.PrepareEvent(producer.Create, 1)
				prod.EXPECT().Send(topic, ev)
				metr.EXPECT().IncCreate()

				resp, err := srv.MultiCreateSurveyV1(ctx, &req)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(resp).ShouldNot(BeNil())
				Expect(resp.GetSurveyIds()).Should(BeEquivalentTo([]uint64{1}))

				Expect(sqlm.ExpectationsWereMet()).ShouldNot(HaveOccurred())
			})
		})

		When("no items passed", func() {
			It("should return error", func() {
				req := desc.MultiCreateSurveyV1Request{}

				resp, err := srv.MultiCreateSurveyV1(ctx, &req)
				Expect(err).Should(HaveOccurred())
				Expect(resp.GetSurveyIds()).Should(BeEmpty())

				Expect(sqlm.ExpectationsWereMet()).ShouldNot(HaveOccurred())
			})
		})
	})

	Describe("DescribeSurveyV1", func() {

		When("specified ID exists", func() {
			It("should return requested item", func() {
				req := &desc.DescribeSurveyV1Request{
					SurveyId: 1,
				}
				data := &desc.Survey{
					Id:     1,
					UserId: 10,
					Link:   "http://api.test/survey/1",
				}

				sqlm.ExpectQuery("SELECT id, user_id, link FROM surveys").
					WithArgs(req.SurveyId).
					WillReturnRows(sqlm.NewRows([]string{"id", "user_id", "link"}).
						AddRow(data.Id, data.UserId, data.Link))

				resp, err := srv.DescribeSurveyV1(ctx, req)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(resp).ShouldNot(BeNil())
				Expect(resp.Survey).Should(BeEquivalentTo(data))

				Expect(sqlm.ExpectationsWereMet()).ShouldNot(HaveOccurred())
			})
		})

		When("item not found", func() {
			It("should return error", func() {
				req := &desc.DescribeSurveyV1Request{
					SurveyId: 2,
				}

				sqlm.ExpectQuery("SELECT id, user_id, link FROM surveys").
					WithArgs(req.SurveyId).
					WillReturnError(sql.ErrNoRows)

				resp, err := srv.DescribeSurveyV1(ctx, req)
				Expect(err).Should(HaveOccurred())
				Expect(resp).Should(BeNil())

				Expect(sqlm.ExpectationsWereMet()).ShouldNot(HaveOccurred())
			})
		})
	})

	Describe("ListSurveysV1", func() {

		When("called with valid arguments", func() {
			It("should return items", func() {
				req := &desc.ListSurveysV1Request{
					Limit:  10,
					Offset: 0,
				}
				data := []*desc.Survey{
					{Id: 1, UserId: 10, Link: "http://api.test/survey/1"},
					{Id: 2, UserId: 20, Link: "http://api.test/survey/2"},
					{Id: 3, UserId: 30, Link: "http://api.test/survey/3"},
				}

				sqlm.ExpectQuery("SELECT id, user_id, link FROM surveys").
					WithArgs(req.Limit, req.Offset).
					WillReturnRows(sqlm.NewRows([]string{"id", "user_id", "link"}).
						AddRow(data[0].Id, data[0].UserId, data[0].Link).
						AddRow(data[1].Id, data[1].UserId, data[1].Link).
						AddRow(data[2].Id, data[2].UserId, data[2].Link))

				resp, err := srv.ListSurveysV1(ctx, req)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(resp).ShouldNot(BeNil())
				Expect(resp.Surveys).Should(BeEquivalentTo(data))

				Expect(sqlm.ExpectationsWereMet()).ShouldNot(HaveOccurred())
			})
		})

		When("no items", func() {
			It("should return empty response", func() {
				req := &desc.ListSurveysV1Request{
					Limit:  10,
					Offset: 0,
				}

				sqlm.ExpectQuery("SELECT id, user_id, link FROM surveys").
					WithArgs(req.Limit, req.Offset).
					WillReturnError(sql.ErrNoRows)

				resp, err := srv.ListSurveysV1(ctx, req)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(resp.Surveys).Should(BeEmpty())

				Expect(sqlm.ExpectationsWereMet()).ShouldNot(HaveOccurred())
			})
		})
	})

	Describe("UpdateSurveyV1", func() {

		When("updating existing survey", func() {
			It("should update data in repo", func() {
				data := &desc.Survey{
					Id:     1,
					UserId: 20,
					Link:   "http://api.test/survey/2",
				}
				req := &desc.UpdateSurveyV1Request{Survey: data}

				sqlm.ExpectExec("UPDATE surveys").
					WithArgs(data.Id, data.UserId, data.Link).
					WillReturnResult(sqlmock.NewResult(0, 1))

				ev := producer.PrepareEvent(producer.Update, 1)
				prod.EXPECT().Send(topic, ev)
				metr.EXPECT().IncUpdate()

				_, err := srv.UpdateSurveyV1(ctx, req)
				Expect(err).ShouldNot(HaveOccurred())

				Expect(sqlm.ExpectationsWereMet()).ShouldNot(HaveOccurred())
			})
		})

		When("updating non-existing survey", func() {
			It("should return error", func() {
				data := &desc.Survey{
					Id:     1,
					UserId: 20,
					Link:   "http://api.test/survey/2",
				}
				req := &desc.UpdateSurveyV1Request{Survey: data}

				sqlm.ExpectExec("UPDATE surveys").
					WithArgs(data.Id, data.UserId, data.Link).
					WillReturnResult(sqlmock.NewResult(0, 0))

				_, err := srv.UpdateSurveyV1(ctx, req)
				Expect(err).Should(HaveOccurred())

				Expect(sqlm.ExpectationsWereMet()).ShouldNot(HaveOccurred())
			})
		})
	})

	Describe("RemoveSurveyV1", func() {

		When("specified ID exists", func() {
			It("should delete requested item", func() {
				req := &desc.RemoveSurveyV1Request{
					SurveyId: 1,
				}

				sqlm.ExpectExec("UPDATE surveys").
					WithArgs(req.SurveyId).
					WillReturnResult(sqlmock.NewResult(0, 1))

				ev := producer.PrepareEvent(producer.Delete, 1)
				prod.EXPECT().Send(topic, ev)
				metr.EXPECT().IncDelete()

				_, err := srv.RemoveSurveyV1(ctx, req)
				Expect(err).ShouldNot(HaveOccurred())

				Expect(sqlm.ExpectationsWereMet()).ShouldNot(HaveOccurred())
			})
		})

		When("item not found", func() {
			It("should return error", func() {
				req := &desc.RemoveSurveyV1Request{
					SurveyId: 2,
				}

				sqlm.ExpectExec("UPDATE surveys").
					WithArgs(req.SurveyId).
					WillReturnResult(sqlmock.NewResult(0, 0))

				_, err := srv.RemoveSurveyV1(ctx, req)
				Expect(err).Should(HaveOccurred())

				Expect(sqlm.ExpectationsWereMet()).ShouldNot(HaveOccurred())
			})
		})
	})
})

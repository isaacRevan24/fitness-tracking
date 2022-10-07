package test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/isaacRevan24/fitness-tracking/handler"
	"github.com/isaacRevan24/fitness-tracking/model"
	"github.com/isaacRevan24/fitness-tracking/repository/mock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLogic(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Handler Test suit")
}

var _ = Describe("Handler tests", func() {

	var (
		mockCtrl       *gomock.Controller
		mockRepository *mock.MockTrackingRepository
		underTest      handler.TrackingHandlerInterface
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockRepository = mock.NewMockTrackingRepository(mockCtrl)
		underTest = handler.NewTrackingHandler()
		handler.Repo = mockRepository
	})

	Context("Save weight entrys", func() {
		It("Should save weight entry and return client Id successfully", func() {
			// Mock
			mockRepository.EXPECT().AddWeightRegister(gomock.Any()).Return("eca98147-8a83-4df0-8812-0668c67b2865", nil)

			// Given
			request := model.AddWeightRegisterReq{Weight: 100.01, CreatedAt: "2021-11-09T11:44:44.797", ClientId: "eca98147-8a83-4df0-8812-0668c67b2865"}

			// When
			response := underTest.AddWeightRegister(request)

			// Then
			Expect(response.Status.Code).Should(Equal(model.SUCCESS_CODE_STATUS))
			Expect(response.Status.Message).Should(Equal(model.SUCCESS_MESSAGE))
			Expect(response.T).ShouldNot(BeNil())
		})

		It("Should fail to add weight register and return error", func() {
			// Mock
			mockRepository.EXPECT().AddWeightRegister(gomock.Any()).Return("", errors.New("error"))

			// Given
			request := model.AddWeightRegisterReq{Weight: 100.01, CreatedAt: "2021-11-09T11:44:44.797", ClientId: "eca98147-8a83-4df0-8812-0668c67b2865"}

			// When
			response := underTest.AddWeightRegister(request)

			// Then
			Expect(response.Status.Code).Should(Equal(model.BAD_REQUEST_ERROR_STATUS))
			Expect(response.Status.Message).Should(Equal(model.ADD_WEIGHT_REGISTER_ERROR))
			Expect(response.T).Should(BeNil())
		})
	})

	Context("Get weight entry", func() {
		It("Should get a weight entry by client and weight id and return weight and date", func() {
			// Mock
			mockResponse := model.GetWeightRegisterRes{Weight: 20.14, CreatedAt: "1969-06-25 00:51:24"}
			mockRepository.EXPECT().GetWeightRegister(gomock.Any(), gomock.Any()).Return(mockResponse, nil)

			// Given
			clientId := "c209cac7-6901-42a1-8e3f-632aecd9911d"
			weightId := "24793e94-eb3c-4580-a6df-ebc09c6eca84"

			// When
			response := underTest.GetWeightRegister(clientId, weightId)

			// Then
			Expect(response.Status.Code).Should(Equal(model.SUCCESS_CODE_STATUS))
			Expect(response.Status.Message).Should(Equal(model.SUCCESS_MESSAGE))
			Expect(response.T).ShouldNot(BeNil())
		})

		It("Should fail to get weight entry and return error status", func() {
			// Mock
			mockRepository.EXPECT().GetWeightRegister(gomock.Any(), gomock.Any()).Return(model.GetWeightRegisterRes{}, errors.New("error"))

			// Given
			clientId := "c209cac7-6901-42a1-8e3f-632aecd9911d"
			weightId := "24793e94-eb3c-4580-a6df-ebc09c6eca84"

			// When
			response := underTest.GetWeightRegister(clientId, weightId)

			// Then
			Expect(response.Status.Code).Should(Equal(model.BAD_REQUEST_ERROR_STATUS))
			Expect(response.Status.Message).Should(Equal(model.GET_WEIGHT_REGISTER_ERROR))
			Expect(response.T).Should(BeNil())
		})

	})

	Context("Update weight entry", func() {
		It("Should successfully update weight entry", func() {
			// Mock
			mockRepository.EXPECT().UpdateWeightRegister(gomock.Any()).Return(nil)

			// Given
			request := model.UpdateWeightRegisterReq{ClientId: "c209cac7-6901-42a1-8e3f-632aecd9911d", WeightTrackId: "24793e94-eb3c-4580-a6df-ebc09c6eca84", Weight: 99.9}

			// When
			response := underTest.UpdateWeightRegister(request)

			// Then
			Expect(response.Status.Code).Should(Equal(model.SUCCESS_CODE_STATUS))
			Expect(response.Status.Message).Should(Equal(model.SUCCESS_MESSAGE))
			Expect(response.T).Should(BeNil())
		})

		It("Should fail to delete wieght entry", func() {
			// Mock
			mockRepository.EXPECT().UpdateWeightRegister(gomock.Any()).Return(errors.New("error"))

			// Given
			request := model.UpdateWeightRegisterReq{ClientId: "c209cac7-6901-42a1-8e3f-632aecd9911d", WeightTrackId: "24793e94-eb3c-4580-a6df-ebc09c6eca84", Weight: 99.9}

			// When
			response := underTest.UpdateWeightRegister(request)

			// Then
			Expect(response.Status.Code).Should(Equal(model.BAD_REQUEST_ERROR_STATUS))
			Expect(response.Status.Message).Should(Equal(model.UPDATE_WEIGHT_REGISTER_ERROR))
			Expect(response.T).Should(BeNil())
		})

	})
})

package test

import (
	"errors"

	"github.com/golang/mock/gomock"
	"github.com/isaacRevan24/fitness-tracking/handler"
	"github.com/isaacRevan24/fitness-tracking/model"
	"github.com/isaacRevan24/fitness-tracking/repository/mock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Goals handler tests", func() {
	var (
		mockCtrl       *gomock.Controller
		mockRepository *mock.MockGoalsRepository
		underTest      handler.GoalsHandlerInterface
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockRepository = mock.NewMockGoalsRepository(mockCtrl)
		underTest = handler.NewGoalsHandler()
		handler.GoalsRepo = mockRepository
	})

	Context("Save goals entrys", func() {
		It("Should save goals entry successfully", func() {
			// Mock
			mockResponse := model.AddWeightGoalsRes{Weight: 99.9, Steps: 13000}
			mockRepository.EXPECT().AddGoalsRegister(gomock.Any()).Return(mockResponse, nil)

			// Given
			request := model.AddWeightGoalsReq{ClientId: "c209cac7-6901-42a1-8e3f-632aecd9911d", Weight: 99.9, Steps: 13000}

			// When
			response := underTest.AddGoalsRegister(request)

			// Then
			Expect(response.Status.Code).Should(Equal(model.SUCCESS_CODE_STATUS))
			Expect(response.Status.Message).Should(Equal(model.SUCCESS_MESSAGE))
			Expect(response.T).ShouldNot(BeNil())
		})

		It("Should fail to save goals entry", func() {
			// Mock
			mockResponse := model.AddWeightGoalsRes{}
			mockRepository.EXPECT().AddGoalsRegister(gomock.Any()).Return(mockResponse, errors.New("error"))

			// Given
			request := model.AddWeightGoalsReq{ClientId: "c209cac7-6901-42a1-8e3f-632aecd9911d", Weight: 99.9, Steps: 13000}

			// When
			response := underTest.AddGoalsRegister(request)

			// Then
			Expect(response.Status.Code).Should(Equal(model.BAD_REQUEST_ERROR_STATUS))
			Expect(response.Status.Message).Should(Equal(model.ADD_GOALS_REGISTER_ERROR))
			Expect(response.T).Should(BeNil())
		})

	})

	Context("Get goals entrys", func() {
		It("Should get goals entry successfully", func() {
			// Mock
			mockResponse := model.GetGoalsRes{Weight: 100}
			mockRepository.EXPECT().GetGoalsRegister(gomock.Any()).Return(mockResponse, nil)

			// Given
			clientId := "c209cac7-6901-42a1-8e3f-632aecd9911d"

			// When
			response := underTest.GetGoalsRegister(clientId)

			// Then
			Expect(response.Status.Code).Should(Equal(model.SUCCESS_CODE_STATUS))
			Expect(response.Status.Message).Should(Equal(model.SUCCESS_MESSAGE))
			Expect(response.T).ShouldNot(BeNil())
		})

		It("Should fail to get goals entry", func() {
			// Mock
			mockRepository.EXPECT().GetGoalsRegister(gomock.Any()).Return(model.GetGoalsRes{}, errors.New("error"))

			// Given
			clientId := "c209cac7-6901-42a1-8e3f-632aecd9911d"

			// When
			response := underTest.GetGoalsRegister(clientId)

			// Then
			Expect(response.Status.Code).Should(Equal(model.BAD_REQUEST_ERROR_STATUS))
			Expect(response.Status.Message).Should(Equal(model.GET_GOALS_REGISTER_ERROR))
			Expect(response.T).Should(BeNil())
		})

	})

	Context("Update goals entrys", func() {
		It("Should update goals entry successfully", func() {
			// Mock
			mockResponse := model.UpdateWeightGoalRes{Weight: 100.2}
			mockRepository.EXPECT().UpdateWeightGoal(gomock.Any()).Return(mockResponse, nil)

			// Given
			request := model.UpdateWeightGoalReq{ClientId: "c209cac7-6901-42a1-8e3f-632aecd9911d", Weight: 100.2}

			// When
			response := underTest.UpdateWeightGoal(request)

			// Then
			Expect(response.Status.Code).Should(Equal(model.SUCCESS_CODE_STATUS))
			Expect(response.Status.Message).Should(Equal(model.SUCCESS_MESSAGE))
			Expect(response.T).ShouldNot(BeNil())
		})

		It("Should fail to get goals entry", func() {
			// Mock
			mockRepository.EXPECT().UpdateWeightGoal(gomock.Any()).Return(model.UpdateWeightGoalRes{}, errors.New("error"))

			// Given
			request := model.UpdateWeightGoalReq{ClientId: "c209cac7-6901-42a1-8e3f-632aecd9911d", Weight: 100.2}

			// When
			response := underTest.UpdateWeightGoal(request)

			// Then
			Expect(response.Status.Code).Should(Equal(model.BAD_REQUEST_ERROR_STATUS))
			Expect(response.Status.Message).Should(Equal(model.UPDATE_WEIGHT_GOAL_REGISTER_ERROR))
			Expect(response.T).Should(BeNil())
		})

	})

})

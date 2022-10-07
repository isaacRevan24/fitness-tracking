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
			mockRepository.EXPECT().AddGoalsRegisters(gomock.Any()).Return(mockResponse, nil)

			// Given
			request := model.AddWeightGoalsReq{ClientId: "c209cac7-6901-42a1-8e3f-632aecd9911d", Weight: 99.9, Steps: 13000}

			// When
			response := underTest.AddGoalsRegisters(request)

			// Then
			Expect(response.Status.Code).Should(Equal(model.SUCCESS_CODE_STATUS))
			Expect(response.Status.Message).Should(Equal(model.SUCCESS_MESSAGE))
			Expect(response.T).ShouldNot(BeNil())
		})

		It("Should fail to save goals entry", func() {
			// Mock
			mockResponse := model.AddWeightGoalsRes{}
			mockRepository.EXPECT().AddGoalsRegisters(gomock.Any()).Return(mockResponse, errors.New("error"))

			// Given
			request := model.AddWeightGoalsReq{ClientId: "c209cac7-6901-42a1-8e3f-632aecd9911d", Weight: 99.9, Steps: 13000}

			// When
			response := underTest.AddGoalsRegisters(request)

			// Then
			Expect(response.Status.Code).Should(Equal(model.BAD_REQUEST_ERROR_STATUS))
			Expect(response.Status.Message).Should(Equal(model.ADD_GOALS_REGISTER_ERROR))
			Expect(response.T).Should(BeNil())
		})

	})

})

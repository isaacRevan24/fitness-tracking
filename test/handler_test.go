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

	It("Should save weight entry and return client Id", func() {
		// Mock
		mockRepository.EXPECT().AddWeightRegister(gomock.Any()).Return("eca98147-8a83-4df0-8812-0668c67b2865", nil)

		// Given
		request := model.AddWeightRegisterReq{Weight: 100.01, CreatedAt: "2021-11-09T11:44:44.797", ClientId: "eca98147-8a83-4df0-8812-0668c67b2865"}

		// When
		response := underTest.AddWeightRegister(request)

		// Then
		Expect(response.Status.Code).Should(Equal(model.SUCCESS_CODE_STATUS))
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
		Expect(response.T).Should(BeNil())
	})

})

package test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/isaacRevan24/fitness-tracking/logic"
	"github.com/isaacRevan24/fitness-tracking/model"
	"github.com/isaacRevan24/fitness-tracking/test/mock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestController(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Logic Test suit")
}

var _ = Describe("Logic tests", func() {

	var (
		mockCtrl       *gomock.Controller
		mockRepository *mock.MockTrackingRepository
		underTest      logic.TrackingLogicInterface
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockRepository = mock.NewMockTrackingRepository(mockCtrl)
		underTest = logic.NewTrackingLogic(mockRepository)
	})

	It("Should save weight entry and return client Id", func() {
		// Mock
		id := uuid.New()
		mockRepository.EXPECT().AddWeightRegister(gomock.Any()).Return(id, nil)

		// Given
		request := model.AddWeightRegisterReq{Weight: 100.01, CreatedAt: "2021-11-09T11:44:44.797", ClientId: uuid.UUID{}}

		// When
		response, err := underTest.AddWeightRegister(request)

		// Then
		Expect(response).Should(Equal(id))
		Expect(err).Should(BeNil())
	})

	It("Should fail to add weight register and return error", func() {
		// Mock
		mockRepository.EXPECT().AddWeightRegister(gomock.Any()).Return(uuid.Nil, errors.New("error"))

		// Given
		request := model.AddWeightRegisterReq{Weight: 100.01, CreatedAt: "2021-11-09T11:44:44.797", ClientId: uuid.UUID{}}

		// When
		response, err := underTest.AddWeightRegister(request)

		// Then
		Expect(response).Should(Equal(uuid.Nil))
		Expect(err).Error()
	})

})

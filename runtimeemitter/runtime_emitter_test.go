package runtimeemitter_test

import (
	"runtime"
	"time"

	"code.cloudfoundry.org/go-loggregator/internal/loggregator_v2"
	"code.cloudfoundry.org/go-loggregator/runtimeemitter"
	"code.cloudfoundry.org/go-loggregator/v2"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RuntimeEmitter", func() {
	It("emits go runtime metrics on an interval", func() {
		v2Client := newSpyV2Client()
		emitter := runtimeemitter.New(v2Client,
			runtimeemitter.WithInterval(10*time.Millisecond),
		)

		go emitter.Run()

		Eventually(v2Client.emitGaugeCalled).Should(BeNumerically(">", 1))
	})

	It("emits expected runtime metrics", func() {
		// Force garbage collection so GC stats aren't empty
		runtime.GC()

		v2Client := newSpyV2Client()
		emitter := runtimeemitter.New(v2Client,
			runtimeemitter.WithInterval(10*time.Millisecond),
		)

		go emitter.Run()

		var env *loggregator_v2.Envelope
		Eventually(v2Client.envelopes).Should(Receive(&env))
		Expect(env.GetGauge()).ToNot(BeNil())

		metrics := env.GetGauge().Metrics
		Expect(metrics["memoryStats.numBytesAllocatedHeap"].Value).To(BeNumerically(">", 0.0))
		Expect(metrics["memoryStats.numBytesAllocatedHeap"].Unit).To(Equal("Bytes"))

		Expect(metrics["memoryStats.numBytesAllocatedStack"].Value).To(BeNumerically(">", 0.0))
		Expect(metrics["memoryStats.numBytesAllocatedStack"].Unit).To(Equal("Bytes"))

		Expect(metrics["numGoRoutines"].Value).To(BeNumerically(">", 0.0))
		Expect(metrics["numGoRoutines"].Unit).To(Equal("Count"))

		Expect(metrics["memoryStats.lastGCPauseTimeNS"].Value).To(BeNumerically(">", 0.0))
		Expect(metrics["memoryStats.lastGCPauseTimeNS"].Unit).To(Equal("ns"))
	})
})

type SpyV2Client struct {
	envelopes chan *loggregator_v2.Envelope
}

func newSpyV2Client() *SpyV2Client {
	return &SpyV2Client{
		envelopes: make(chan *loggregator_v2.Envelope, 100),
	}
}

func (s *SpyV2Client) emitGaugeCalled() int64 {
	return int64(len(s.envelopes))
}

func (s *SpyV2Client) EmitGauge(opts ...v2.EmitGaugeOption) {
	env := &loggregator_v2.Envelope{
		Message: &loggregator_v2.Envelope_Gauge{
			Gauge: &loggregator_v2.Gauge{
				Metrics: make(map[string]*loggregator_v2.GaugeValue),
			},
		},
	}

	for _, o := range opts {
		o(env)
	}

	s.envelopes <- env
}
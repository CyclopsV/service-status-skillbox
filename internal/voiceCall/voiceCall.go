package voiceCall

import (
	"github.com/ferdypruis/iso3166"
	"golang.org/x/exp/slices"
	"strconv"
	"strings"
)

var allowedProviders = []string{"TransparentCalls", "E-Voice", "JustPhone"}

type VoiceCall struct {
	country         string
	bandwidth       int
	avgRespTime     int
	provider        string
	conStability    float32
	clean           int
	avgCallDuration int
	unknownField    int // Так 7 или 8 полей?
}

func New(country, provider string, bandwidth, avgRespTime, clean, avgCallDuration, unknownField int, conStability float32) *VoiceCall {
	if _, err := iso3166.FromAlpha2(country); err != nil {
		return nil
	}
	if bandwidth < 0 && bandwidth > 100 {
		return nil
	}
	if !slices.Contains(allowedProviders, provider) {
		return nil
	}
	return &VoiceCall{
		country:         country,
		bandwidth:       bandwidth,
		avgRespTime:     avgRespTime,
		provider:        provider,
		conStability:    conStability,
		clean:           clean,
		avgCallDuration: avgCallDuration,
		unknownField:    unknownField,
	}
}

func FromSTR(str string) *VoiceCall {
	listStr := strings.Split(str, ";")
	if len(listStr) < 8 {
		return nil
	}
	bandwidth, err := strconv.Atoi(listStr[1])
	if err != nil {
		return nil
	}
	avgRespTime, err := strconv.Atoi(listStr[2])
	if err != nil {
		return nil
	}
	conStability, err := strconv.ParseFloat(listStr[4], 32)
	if err != nil {
		return nil
	}
	clean, err := strconv.Atoi(listStr[5])
	if err != nil {
		return nil
	}
	avgCallDuration, err := strconv.Atoi(listStr[6])
	if err != nil {
		return nil
	}
	unknownField, err := strconv.Atoi(listStr[7])
	if err != nil {
		return nil
	}
	return New(listStr[0], listStr[3], bandwidth, avgRespTime, clean, avgCallDuration, unknownField, float32(conStability))
}

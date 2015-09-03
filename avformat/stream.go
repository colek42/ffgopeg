/*
	AVStream
*/
package avformat

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"

//Rational av_stream_get_r_frame_rate (const Stream *s)
func (s *Stream) AvStreamGetRFrameRate() Rational {
	return (Rational)(C.av_stream_get_r_frame_rate((*C.struct_AVStream)(s)))
}

//void av_stream_set_r_frame_rate (Stream *s, Rational r)
func (s *Stream) AvStreamSetRFrameRate(r Rational) {
	C.av_stream_set_r_frame_rate((*C.struct_AVStream)(s), (C.struct_AVRational)(r))
}

//struct CodecParserContext * av_stream_get_parser (const Stream *s)
func (s *Stream) AvStreamGetParser() *CodecParserContext {
	return (*CodecParserContext)(C.av_stream_get_parser((*C.struct_AVStream)(s)))
}

// //char * av_stream_get_recommended_encoder_configuration (const Stream *s)
// func (s *Stream) AvStreamGetRecommendedEncoderConfiguration() string {
// 	return C.GoString(C.av_stream_get_recommended_encoder_configuration((*C.struct_AVStream)(s)))
// }

// //void av_stream_set_recommended_encoder_configuration (Stream *s, char *configuration)
// func (s *Stream) AvStreamSetRecommendedEncoderConfiguration( c string) {
// 	C.av_stream_set_recommended_encoder_configuration((*C.struct_AVStream)(s), C.CString(c))
// }

//int64_t av_stream_get_end_pts (const Stream *st)
//Returns the pts of the last muxed packet + its duration.
func (s *Stream) AvStreamGetEndPts() int64 {
	return int64(C.av_stream_get_end_pts((*C.struct_AVStream)(s)))
}
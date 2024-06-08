// suppose we are  upgrading existing media player with feature like play_audio_with_caption  and play360video
//  but our old interface lets call(vlc_box) only support function
// play_video and play_audio now we want that our new feature also compatible with older interface.

package main

import "fmt"

type user struct{} //user who want to play audio or video

// Below are the types of new features
type newvideofeature struct{} // play360video
type newaudiofeature struct{} //play_audio_with_caption feature

// Below are already implemented features (old features)
type old_audio_video_player struct{}

// Interface that support old features
type vlc_box interface {
	play_video()
	play_audio()
}

// Below are existing features implementation
func (old old_audio_video_player) play_video() {
	fmt.Println("succesfully played video")
}
func (old old_audio_video_player) play_audio() {
	fmt.Println("succesfully played audio")
}

// Now we need something so that our new features should be compatible with our old interface
// to fulfill this requirement we create adapter type
type vlc_adapter struct {
	videofeature newvideofeature
	audiofeature newaudiofeature
}

// This is adapter for newvideo feature
func (adapter vlc_adapter) play_video() {
	fmt.Println("This is where adapter helps to use newfeature i.e  newvideofeature even with  old interface")
	adapter.videofeature.play360video()

}

// this is implementation of newvideo feature
func (new newvideofeature) play360video() {
	fmt.Println("successfully using 360 video feature")
}

// This is adapter for newaudio feature
func (adapter vlc_adapter) play_audio() {
	fmt.Println("This is where adapter helps to use newfeature i.e  newaudiofeature even with  old interface")
	adapter.audiofeature.play_audio_with_caption()

}

// this is implementation of newaudio feature
func (new newaudiofeature) play_audio_with_caption() {
	fmt.Println("successfully using audio with caption feature")
}

// utility function to use new video service
func (u user) NewVideoService(player vlc_box) {
	player.play_video()

}

// utility function to use new audio service
func (u user) NewAudioService(player vlc_box) {
	player.play_audio()

}

func main() {
	User := user{}
	videofeature := newvideofeature{}
	audiofeature := newaudiofeature{}
	adapter := vlc_adapter{videofeature, audiofeature}

	User.NewVideoService(adapter)
	User.NewAudioService(adapter)

}

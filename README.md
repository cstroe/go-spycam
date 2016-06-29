# Capture frames from a Webcam

* Selects the webcam's MJPEG format at the highest supported resolution.
* Saves frames to a directory.

After you capture frames, [Stackexchange](http://video.stackexchange.com/a/7913) says you can encode them into MKV:

```
ffmpeg -framerate 15 -i frame_%05d.jpeg -codec copy output.mkv
```

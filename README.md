# Capture frames from a Webcam

* Selects the webcam's MJPEG format at the highest supported resolution.
* Saves frames to a directory.

After you capture frames, you can encode into MKV:

```
ffmpeg -framerate 15 -i frame_%05d.jpeg -codec copy output.mkv
```

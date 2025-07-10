# ID3 Tag Editor

I will find a better title later

## Testing

### Configure

Install some prettier stuff for the test output: `go install gotest.tools/gotestsum@latest` or install using brew: `brew install gotestsum`.

### Test the size of the ID3 tag

You need to get the actual size using alternate means. I opted to use `ffmpeg`:

```bash
ffmpeg -i sample.mp3 -v debug 2>&1 | grep id3v2 | awk '{print $4}' | awk -F ':' '{print $2}'
```

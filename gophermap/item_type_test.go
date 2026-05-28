package gophermap

import (
	"path/filepath"
	"testing"
)

func TestNewItemTypeFromPath(t *testing.T) {
	tests := []struct {
		path     string
		expected byte
	}{
		{path: "/a/b/", expected: ItemTypeGopherMenu},
		{path: "/", expected: ItemTypeGopherMenu},
		{path: "/a/b.txt", expected: ItemTypeTextFile},
		{path: "/document.txt", expected: ItemTypeTextFile},
		{path: "/a/b.gif", expected: ItemTypeGIFFile},
		{path: "/image.gif", expected: ItemTypeGIFFile},
		{path: "/a/b.html", expected: ItemTypeHTML},
		{path: "/style.css", expected: ItemTypeHTML},
		{path: "/script.js", expected: ItemTypeHTML},
		{path: "a/b.bin", expected: ItemTypeBinaryFile},
		{path: "/disk.img", expected: ItemTypeBinaryFile},
		{path: "/system.iso", expected: ItemTypeBinaryFile},
		{path: "a/b.hqx", expected: ItemTypeBinHexEncodedFile},
		{path: "/a/b.ph", expected: ItemTypeCCSONameserver},
		{path: "/a/b.exe", expected: ItemTypeDOSBinary},
		{path: "/program.com", expected: ItemTypeDOSBinary},
		{path: "/library.dll", expected: ItemTypeDOSBinary},
		{path: "/driver.sys", expected: ItemTypeDOSBinary},
		{path: "/a/b.uu", expected: ItemTypeUNIXUuencodedFile},
		{path: "/file.uue", expected: ItemTypeUNIXUuencodedFile},
		{path: "/audio.mp3", expected: ItemTypeSoundFile},
		{path: "/sound.wav", expected: ItemTypeSoundFile},
		{path: "/music.flac", expected: ItemTypeSoundFile},
		{path: "song.ogg", expected: ItemTypeSoundFile},
		{path: "track.aac", expected: ItemTypeSoundFile},
		{path: "podcast.m4a", expected: ItemTypeSoundFile},
		{path: "voice.opus", expected: ItemTypeSoundFile},
		{path: "recording.wma", expected: ItemTypeSoundFile},
		{path: "audio.webm", expected: ItemTypeSoundFile},
		{path: "image.png", expected: ItemTypeOtherImageFile},
		{path: "photo.jpg", expected: ItemTypeOtherImageFile},
		{path: "picture.jpeg", expected: ItemTypeOtherImageFile},
		{path: "icon.svg", expected: ItemTypeOtherImageFile},
		{path: "graphic.webp", expected: ItemTypeOtherImageFile},
		{path: "bitmap.bmp", expected: ItemTypeOtherImageFile},
		{path: "favicon.ico", expected: ItemTypeOtherImageFile},
		{path: "photo.tiff", expected: ItemTypeOtherImageFile},
		{path: "animation.apng", expected: ItemTypeOtherImageFile},
		{path: "modern.avif", expected: ItemTypeOtherImageFile},
		{path: "a/b.unknown", expected: ItemTypeGopherMenu},
		{path: "file.xyz", expected: ItemTypeGopherMenu},
		{path: "noextension", expected: ItemTypeGopherMenu},
	}

	for _, test := range tests {
		extension := filepath.Ext(test.path)

		itemType := NewItemTypeFromFileExtension(extension)
		if itemType != test.expected {
			t.Fatalf(
				"'%c' is not the right item type for the path: '%s' (expected: '%c')",
				itemType,
				test.path,
				test.expected,
			)
		}
	}
}

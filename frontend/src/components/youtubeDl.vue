<template>
  <div id="select-value-extension-container">
    <select :value="selectedExt" id="select-value-extension">
      <option id="defaultext" value="defaultext">Default</option>
      <optgroup label="Video">
        <option id="mp4" value="mp4">mp4</option>
        <option id="webm" value="webm">webm</option>
      </optgroup>
      <optgroup label="Audio">
        <option id="m4a" value="m4a">m4a</option>
        <option id="mp3" value="mp3">mp3</option>
        <option id="vorbis" value="vorbis">vorbis</option>
      </optgroup>
    </select>
  </div>
</template>
<script>
export default {
    name:"youtubeDl",
}
</script>

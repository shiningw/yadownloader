<template>
  <form class="main-form" id="nc-vue-unified-form" :action="path">
    <div class="options-group">
      <div
        class="magnet-link http-link option-buttons active-button"
        data-type="aria2"
        @click.prevent="configure"
      >
        HTTP/MAGNET
      </div>
      <div
        v-if="ytd"
        data-type="ytd"
        class="youtube-dl-link option-buttons"
        @click.prevent="configure"
      >
        Youtube-dl
      </div>
      <div
        v-if="search"
        data-type="search"
        class="search-torrents option-buttons"
        @click.prevent="configure"
      >
        {{ searchLabel }}
      </div>
    </div>
    <div class="action-group">
      <div class="download-input-container" v-if="inputType === 'download'">
        <textInput :placeholder="placeholder" :dataType="downloader"></textInput>
        <div class="download-controls-container">
          <youtube-dl v-if="selectOptions"></youtube-dl>
          <actionButton className="download-button" @clicked="download"></actionButton>
          <uploadFile
            v-if="downloader === 'aria2'"
            @uploadfile="uploadFile"
            :path="uris.upload_url"
          ></uploadFile>
        </div>
      </div>
      <searchInput
        v-else
        @search="search"
        @optionSelected="optionCallback"
        :selectOptions="searchOptions"
      ></searchInput>
    </div>
  </form>
</template>
<script>
import textInput from "./textInput.vue";
import searchInput from "./searchInput.vue";
import actionButton from "./actionButton.vue";
import uploadFile from "./uploadFile.vue";
import YoutubeDl from "./youtubeDl.vue";

export default {
  inject: ["settings", "search_sites"],
  data() {
    return {
      checkedValue: false,
      path: this.uris.aria2_url,
      dlPath: this.settings.ncd_downloader_dir,
      inputType: "download",
      selectOptions: false,
      downloader: "aria2",
      placeholder: "Paste your http/magnet link here",
      searchLabel: "Search Torrents",
      searchOptions: this.search_sites ? this.search_sites : this.noOptions(),
      selectedExt: "defaultext",
      ytd: true,
      search: false,
    };
  },
  components: {
    textInput,
    actionButton,
    searchInput,
    uploadFile,
    YoutubeDl,
  },
  created() {},
  computed: {},
  methods: {
    configure(event) {
      let element = event.target;
      let nodeList = element.parentElement.querySelectorAll(".option-buttons");
      nodeList.forEach((node) => {
        node.classList.remove("active-button");
      });
      element.classList.toggle("active-button");
      this.downloader = element.getAttribute("data-type");
      if (this.downloader === "aria2") {
        this.path = this.uris.aria2_url;
      } else if (this.downloader === "ytd") {
        this.placeholder = "Paste your video link here";
        this.path = this.uris.ytd_url;
      } else {
        this.path = this.uris.search_url;
      }
      this.selectOptions = this.downloader === "youtube-dl" ? true : false;
      this.inputType = this.downloader === "search" ? "search" : "download";
    },
    download(event) {
      this.$emit("download", event);
    },
    search(event, vm) {
      this.$emit("search", event, vm);
    },
    uploadFile(event, vm) {
      this.$emit("uploadfile", event, vm);
    },
    optionCallback(option) {
      if (option.label.toLowerCase() == "music") {
        this.searchLabel = "Search Music";
      } else {
        this.searchLabel = "Search Torrents";
      }
    },
    noOptions() {
      return [{ name: "nooptions", label: "No Options" }];
    },
  },
  mounted() {},
  name: "mainForm",
  props: {
    uris: Object,
    uri: String,
  },
};
</script>
<style lang="scss">
html {
  box-sizing: border-box;
}
*,
*::before,
*::after {
  box-sizing: inherit;
}
#nc-vue-unified-form {
  display: flex;
  width: 100%;
  height: $column-height;
  font-size: medium;
  .action-group {
    width: 100%;
  }
  .options-group,
  .action-group > div {
    display: flex;
    width: auto;
    height: 100%;
    position: relative;
  }
  .options-group > .option-buttons {
    margin: 0;
    padding: 10px;
    outline: 0;
    font-weight: bold;
    font-size: 13px;
    font-family: inherit;
    vertical-align: baseline;
    cursor: pointer;
    white-space: nowrap;
    min-height: 34px;
    width: auto;
  }
  .active-button {
    border: 2px #9a5c8b solid;
  }

  .action-group {
    flex: 2;
    & > div {
      border: 1px solid #565687;
      & > div,
      & > select {
        height: 100%;
        display: flex;
        padding: 0px;
        margin: 0px;
      }
      & > div[class$="-controls-container"] {
        display: flex;
        & div,
        & select {
          height: 100%;
          color: #181616;
          font-size: medium;
          background-color: #bdbdcf;
        }
      }
    }
  }

  .selectOptions {
    border-radius: 0%;
  }
  .download-button,
  .search-button {
    height: $column-height;
    .btn-primary {
      color: #fff;
      background-color: #2d3f59;
      border-color: #1e324f;
      border-radius: 0%;
    }
    .btn-primary:hover {
      background-color: #191a16;
    }
  }

  .magnet-link,
  .choose-file {
    background-color: #a0a0ae;
    border-radius: 15px 0px 0px 15px;
  }

  .youtube-dl-link {
    background-color: #b8b8ca;
  }
  .search-torrents {
    background-color: #d0d0e0;
  }

  .search-torrents,
  .youtube-dl-link,
  .magnet-link,
  .choose-file {
    color: #181616;
  }
  .selectOptions {
    background-color: #c4c4d9;
    padding: 5px 1px;
  }
  input,
  select,
  button {
    margin: 0px;
    border: 0px;
    padding: 10px;
    height: 100%;
  }
  button {
    white-space: nowrap;
  }
}
@media only screen and (max-width: 1024px) {
  #nc-vue-unified-form {
    display: flex;
    flex-flow: column;
    row-gap: 10px;
    height: $column-height * 3 + 10;
    margin-left: #{$menu-toggle-width};

    .options-group,
    .action-group > div {
      display: flex;
      width: 100%;
      height: $column-height;
    }

    .action-group > div {
      border: 0px;
      flex-flow: column nowrap;
      & > div {
        margin: 5px 1px;
      }
      & > div[class$="-controls-container"] {
        display: flex;
        justify-content: center;
      }
    }
    .options-group {
      & > button {
        width: calc(100% / 3);
      }
    }
  }
}
</style>

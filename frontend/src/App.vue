<template>
  <div class="app-wrapper">
    <AppNavigation :queues="queues"></AppNavigation>
    <div id="app-content">
      <section>
        <div
          id="app-navigation-toggle"
          class="icon-menu"
          style="display: none"
          tabindex="0"
        ></div>
      </section>
      <section v-if="display.download" class="form-section" id="form-section">
        <mainForm :uris="uris" @download="download" @uploadfile="uploadFile"></mainForm>
      </section>
      <section id="downloader-table-wrapper"></section>
    </div>
  </div>
</template>

<script>
import mainForm from "./components/mainForm.vue";
import AppNavigation from "./components/appNavigation.vue";
import helper from "./utils/helper";

export default {
  //inject: ["settings"],
  provide() {
    return {
      search_sites: [
        { name: "pirateby", label: "Piratebay" },
        { name: "p", label: "pb" },
      ],
      settings: helper.getServerSettings(),
    };
  },
  data() {
    return {
      display: { download: true, search: false },
      uris: {
        ytd_url: "/ytd/action/download",
        aria2_url: "/aria2/action/download",
        search_url: "/apps/ncdownloader/search",
        upload_url: "/aria2/upload",
      },
      queues: [
        {
          id: "active",
          label: "Downloading",
          path: "/aria2/data/active",
        },
        {
          id: "failed",
          label: "Failed",
          path: "/aria2/data/failed",
        },
        {
          id: "waiting",
          label: "Waiting",
          path: "/aria2/data/waiting",
        },
        {
          id: "complete",
          label: "Complete",
          path: "/aria2/data/complete",
        },
        {
          id: "ytd",
          label: "Youtube",
          path: "/ytd/downloads",
        },
      ],
    };
  },
  methods: {
    download(event) {
      let element = event.target;
      let formWrapper = element.closest("form");
      let formData = helper.getData(formWrapper);
      let inputValue = formData["text-input-value"].trim();
      if (!helper.isURL(inputValue) && !helper.isMagnetURI(inputValue)) {
        helper.error(inputValue + " is Invalid");
        return;
      }
      let url = helper.generateUrl(formWrapper.getAttribute("action"));
      console.log(formData);
      helper
        .httpClient(url)
        .setData(formData)
        .setHandler(function (data) {
          if (data.error) {
            helper.error(data.error);
          } else if (data.data) {
            helper.info(data.data);
          }
        })
        .send();
    },
    uploadFile(event, vm) {
      let element = event.target;
      const files = element.files || event.dataTransfer.files;
      if (files) {
        let formWrapper = element.closest("form");
        let url = formWrapper.getAttribute("action");
        url = helper.generateUrl(url);
        return helper
          .httpClient(url)
          .setHandler(function (data) {
            if (data.error) {
              helper.error(data.error);
            } else if (data.data) {
              helper.info(data.data);
            }
          })
          .upload(files[0]);
      }
      return false;
    },
  },
  mounted() {
    let d = document.getElementById("app-settings-data");
    let j = JSON.parse(d.getAttribute("data-settings"));
  },
  name: "mainApp",
  components: {
    mainForm,
    AppNavigation,
  },
};
</script>

<style lang="scss">
//@import "css/variables.scss";

#app-content-wrapper {
  .ncdownloader-form-wrapper {
    position: relative;
    width: 100%;
    top: 0;
    left: 0;
  }
  .ncdownloader-form-wrapper.top-left {
    width: 100%;
    top: 0;
    left: 0;
  }

  .form-section {
    width: 100%;
    display: flex;
    flex-flow: column;
    gap: 1.2em;
  }
}

@media only screen and (max-width: 1024px) {
  #app-content-wrapper {
    #ncdownloader-form-wrapper {
      position: relative;
      margin: 2px;
    }
  }
}
</style>

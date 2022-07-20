<template>
  <div id="app-settings">
    <div id="app-settings-header">
      <button
        name="app settings"
        class="settings-button"
        data-apps-slide-toggle="#app-settings-content"
      >
        Settings
      </button>
    </div>
    <div id="app-settings-content">
      <toggleButton
        :disabledText="errorText"
        :enabledText="errorText"
        :defaultStatus="toggleStatus"
        @changed="toggle"
        name="disable_polling"
      ></toggleButton>
    </div>
  </div>
</template>
<script>
//import { registerAppsSlideToggle } from "../lib/appSidebar";
import Slide from "../lib/slide";
import toggleButton from "./toggleButton";
//registerAppsSlideToggle();

export default {
  name: "Sidebar",
  data () {
    return {
      errorText: "disable polling",
      toggleStatus: false,
    };
  },
   methods: {
    toggle(event) {
      this.$emit("toggled", event);
    },
  },
  mounted() {
    Slide.create();
  },
  components: {
    toggleButton,
  },
};
</script>
<style lang="scss">
#app-settings-header .settings-button {
  display: block;
  height: 44px;
  width: 100%;
  padding: 0;
  margin: 0;
  background-image: url($icon-settings);
  background-position: 14px center;
  background-repeat: no-repeat;
  box-shadow: none;
  border: 0;
  border-radius: 0;
  text-align: left;
  padding-left: 44px;
  font-weight: normal;
  font-size: 100%;
  opacity: 0.8;
}
#app-settings-content {
  display: none;
  padding: 10px;
  max-height: 300px;
  overflow-y: auto;
  box-sizing: border-box;
}
</style>

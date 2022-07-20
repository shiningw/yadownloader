<template>
  <div id="app-navigation">
    <iconButton classes="icon-add" :disabled="true"
      >Download</iconButton
    >
    <iconButton path="/aria2/start" @clicked.prevent="startAria2" classes="icon-power">{{
      aria2IsRunning ? "Stop Aria2" : "Start Aria2"
    }}</iconButton>
    <ul>
      <li v-for="(queue, index) in queues" :key="index" class="download-queue">
        <div class="entry-bullet"></div>
        <a
          role="button"
          tabindex="0"
          :path="queue.path"
          :id="queue.id + '-queue'"
          @click.prevent="polling($event, queue.path)"
        >
          {{ queue.label }}</a
        >
        <div class="entry-utils">
          <ul>
            <li class="entry-utils-counter">
              <div class="number counter-">{{ counters[queue.id] }}</div>
            </li>
          </ul>
        </div>
      </li>
    </ul>
    <Sidebar @toggled="disablePolling"></Sidebar>
  </div>
</template>

<script>
import Sidebar from "./sidebar.vue";
import helper from "../utils/helper";
import iconButton from "./iconButton.vue";
export default {
  inject: ["settings"],
  data() {
    //const settings = helper.getServerSettings();
    let counters = this.settings.counters;
    return {
      aria2IsRunning: false,
      counters: {
        active: counters.active,
        failed: counters.failed,
        waiting: counters.waiting,
        complete: counters.complete,
        ytd: counters.ytd,
      },
    };
  },
  components: {
    Sidebar,
    iconButton,
  },
  methods: {
    executeCallback(event, callback) {
      callback(event);
    },
    polling(event, path) {
      helper.polling(2000, path);
    },
    disablePolling() {
      helper.disablePolling();
    },
    startAria2: async function (event, vm) {
      let data = await helper.aria2IsRunning();
      if (data.status) {
        const resp = await helper.stopAria2();
        helper.sleep(2000).then(() => {
          vm.$data.loading = false;
          if (resp.status) {
            this.aria2IsRunning = false;
          }
        });
      } else {
        helper.startAria2().then((resp) => {
          vm.$data.loading = false;
          if (resp.error) {
            helper.error(resp.error);
          }
          if (resp.status) {
            this.aria2IsRunning = true;
          }
        });
      }
    },
  },
  mounted() {
    this.aria2IsRunning = this.settings.aria2.status;
  },
  name: "appNavigation",
  props: {
    queues: Array,
  },
};
</script>
<style lang="scss">
#app-navigation {
  width: $navigation-width;
  display: flex;
  height: 100%;
  flex-direction: column;
  overflow-y: auto;
  overflow-x: hidden;
  box-sizing: border-box;
  position: fixed;
  left: 0;
  z-index: 500;
  background-color: $color-main-background;
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
  border-right: 1px solid $color-border;

  & > ul > li,
  ul > li > ul > li {
    position: relative;
    width: 100%;
    box-sizing: border-box;
  }
  & > ul > li {
    display: inline-flex;
    flex-wrap: wrap;
    order: 1;
    flex-shrink: 0;
  }
  ul > li > .entry-bullet {
    position: absolute;
    display: block;
    margin: 16px;
    width: 12px;
    height: 12px;
    border: none;
    border-radius: 50%;
    cursor: pointer;
    transition: background 100ms ease-in-out;
  }

  .entry-bullet {
    background-color: #5959c1;
  }

  .entry-utils {
    flex: 0 1 auto;
    ul {
      display: flex !important;
      align-items: center;
      justify-content: flex-end;
      li {
        width: 44px !important;
        height: 44px;
        position: relative;
      }

      .entry-utils-counter {
        overflow: hidden;
        text-align: right;
        font-size: 9pt;
        line-height: 44px;
        padding: 0 12px;
      }
    }
  }

  ul {
    position: relative;
    height: 100%;
    width: inherit;
    overflow-x: hidden;
    overflow-y: auto;
    box-sizing: border-box;
    display: flex;
    flex-direction: column;
  }

  ul > li > a {
    padding: 0 12px 0 44px;
    background-size: 16px 16px;
    background-position: 14px center;
    background-repeat: no-repeat;
    display: block;
    justify-content: space-between;
    line-height: 44px;
    min-height: 44px;
    overflow: hidden;
    box-sizing: border-box;
    white-space: nowrap;
    text-overflow: ellipsis;
    color: var(--color-main-text);
    opacity: 0.8;
    flex: 1 1 0px;
    z-index: 100;
  }
  .entry-bullet + a {
    background: transparent !important;
  }
}
@media only screen and (max-width: 1024px) {
  .snapjs-left #app-navigation {
    transform: translateX(0);
  }
  #app-navigation {
    transform: translateX(-$navigation-width);
  }
  #app-navigation-toggle {
    position: fixed;
    display: inline-block !important;
    left: 0;
    width: 44px;
    height: 44px;
    z-index: 1050;
    cursor: pointer;
    opacity: 0.6;
  }
}
</style>

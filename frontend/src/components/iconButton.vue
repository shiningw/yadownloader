<template>
  <div class="navigation-btn">
    <button
      class="button"
      :class="[classes, loading ? 'is-loading' : '']"
      :path="path"
      @click="handler"
    >
      <slot>Action button</slot>
    </button>
  </div>
</template>
<script>
export default {
  name: "iconButton",
  data() {
    return {
      loading: false,
    };
  },
  methods: {
    handler(event) {
      if (this.disabled) {
        return;
      }
      this.loading = true;
      this.$emit("clicked", event, this);
    },
  },
  props: {
    path: String,
    classes: String,
    disabled: {
      type: Boolean,
      default: false,
    },
  },
};
</script>
<style lang="scss"></style>

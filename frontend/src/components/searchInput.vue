<template>
  <div class="search-input" id="nc-vue-search-input">
    <textInput :placeholder="placeholder" dataType="search"></textInput>
    <div class="search-controls-container">
      <div id="select-value-search-container">
        <select :value="selected" @change="selectHandler" id="select-value-search">
          <option
            v-for="(option, index) in selectOptions"
            v-bind:key="index"
            :value="option.name"
          >
            {{ option.label }}
          </option>
        </select>
      </div>
      <actionButton className="search-button" :enableLoading="true" @clicked="search"
        >Search</actionButton
      >
    </div>
  </div>
</template>
<script>
import textInput from "./textInput.vue";
import actionButton from "./actionButton.vue";

export default {
  data() {
    return {
      placeholder: "Enter keyword to search",
      selected: "TPB",
    };
  },
  components: {
    textInput,
    actionButton,
  },
  methods: {
    search(event, btnVm) {
      this.$emit("search", event, btnVm);
    },
    selectHandler(event) {
      const data = {};
      const element = event.target;
      data.key = element.value;
      data.label = element.options[element.selectedIndex].text;
      this.$emit("optionSelected", data);
    },
  },
  name: "searchInput",
  props: {
    selectOptions: Array,
  },
};
</script>
<style scoped lang="scss">
</style>

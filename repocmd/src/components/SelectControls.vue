<template>
  <div class="select-controls">
    <div class="action-group">
      <h4>Apply an action</h4>
      <dropdown class="action-dropdown" placeholder="Select an action..." :options="dropdownOpts" :selected="selected" @updateOption="select"></dropdown>
      <a class="action-btn" href="#">I understand, do it</a>
    </div>
    <p v-if="selected == 'archive'">
      <i data-feather="info"></i> Archiving can only be reversed through the GitHub website.
    </p>
  </div>
</template>

<script>
import feather from 'feather-icons'
import dropdown from 'vue-dropdowns'
export default {
  components: {
    dropdown
  },
  data () {
    return {
      selected: '',
      dropdownOpts: [{name: 'Archive', value: 'archive'}]
    }
  },
  mounted () {
    feather.replace()
  },
  methods: {
    select (payload) {
      this.selected = payload.value
    }
  }
}
</script>

<style lang="scss">
  optgroup {
    font-size: 20px;
  }

  .feather {
    height: 16px;
    width: 16px;
  }

  .action-group {
    margin-top: 20px;
    text-align: center;

    .action-dropdown {
      margin: 0 !important;
      margin-right: 10px !important;
      vertical-align: unset !important;
      // Dropdown overries
      .dropdown-toggle {
        border: 1px solid #000 !important;
        padding: 12px 20px !important;
        // background: transparent;
        background-image: none !important;
        color: var(--black) !important;
        list-style: none;
      }
      .dropdown-menu > li {
        & > a {
          font-size: 26px;
          &:hover {
            color: #000 !important;
          }
        }
      }
    }
    .action-btn {
      border: 1px solid red;
      font-weight: normal;
      font-size: 16px;
      color: #FF2A2A;
      text-decoration: none;
      padding: 12px 20px;
      display: inline-block;
    }
  }
</style>

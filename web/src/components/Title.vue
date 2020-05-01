<template>
  <div style="text-align: left; margin-top: 20px; font-weight: bold">
  {{this.title}}
  </div>
</template>

<script>
export default {
  name: 'Title',
  props: {
    propTitle: String
  },
  data: function () {
    return {
      title: this.propTitle
    }
  }
}
</script>

<style scoped>

</style>

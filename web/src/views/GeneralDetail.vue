<template>
  <div class="generalDetail">
    <Detail :prop-board-name="boardName"></Detail>
  </div>
</template>

<script>
export default {
  name: 'GeneralDetail',
  data: function () {
    return {
      boardName: 'general'
    }
  }
}
</script>

<style scoped>

</style>

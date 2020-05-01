<template>
  <div class="generalDetail">
    <NavMenu :prop-index="activeIndex"></NavMenu>
    <Detail :prop-board-name="boardName"></Detail>
  </div>
</template>

<script>
export default {
  name: 'GeneralDetail',
  data: function () {
    return {
      activeIndex: '2-1',
      boardName: 'general'
    }
  }
}
</script>

<style scoped>

</style>

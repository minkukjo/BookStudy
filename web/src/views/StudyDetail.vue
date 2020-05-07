<template>
  <div class="studyDetail">
    <Detail :prop-board-name="boardName"></Detail>
  </div>
</template>

<script>
export default {
  name: 'StudyDetail',
  data: function () {
    return {
      boardName: 'study'
    }
  }
}
</script>

<style scoped>

</style>

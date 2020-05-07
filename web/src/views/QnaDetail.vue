<template>
  <div class="qnaDetail">
    <Detail :prop-board-name="boardName"></Detail>
  </div>
</template>

<script>
export default {
  name: 'QnaDetail',
  data: function () {
    return {
      boardName: 'qna'
    }
  }
}
</script>

<style scoped>

</style>

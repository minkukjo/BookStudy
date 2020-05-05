<template>
  <div class="qna">
    <NavMenu :prop-index="activeIndex"></NavMenu>
    <Title :prop-title="title"></Title>
    <Board v-if="posts" :prop-posts="posts" :prop-board-name="boardName"></Board>
  </div>
</template>

<script>
export default {
  name: 'Qna',
  data: function () {
    return {
      activeIndex: '2-3',
      title: '질문 게시판',
      boardName: 'qna',
      posts: []
    }
  },
  created () {
    this.getAllPosts()
  },
  methods: {
    getAllPosts: function () {
      this.$http.get(this.$http_url + '/api/posts?kind=qna')
        .then((response) => {
          this.posts = response.data
        })
    }
  }
}
</script>

<style scoped>

</style>

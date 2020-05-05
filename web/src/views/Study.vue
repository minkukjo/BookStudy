<template>
  <div class="study">
    <NavMenu :prop-index="activeIndex"></NavMenu>
    <Title :prop-title="title"></Title>
    <Board v-if="posts" :prop-posts="posts" :prop-board-name="boardName"></Board>
  </div>
</template>

<script>
export default {
  name: 'Study',
  data: function () {
    return {
      activeIndex: '2-2',
      title: '스터디 모집',
      boardName: 'study',
      posts: []
    }
  },
  created () {
    this.getAllPosts()
  },
  methods: {
    getAllPosts: function () {
      this.$http.get(this.$http_url + '/api/posts?kind=study')
        .then((response) => {
          this.posts = response.data
        })
    }
  }
}
</script>

<style scoped>

</style>

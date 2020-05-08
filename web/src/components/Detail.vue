<template>
  <div class="content">
    <div class="head">
      <i class="el-icon-back" @click="goBack"></i>
      <i class="el-icon-delete"></i>
      <div>
        <div class="name">
          {{post.name}}
        </div>
        {{post.date}}
      </div>
    </div>
    <div class="content-main">
      <div class="content-body">
        <div class="title">
          {{post.title}}
        </div>
        <div class="text">
          {{post.text}}
        </div>
      </div>
      <div class="content-function">
        추천버튼
      </div>
    </div>
  </div>
</template>

<script>

export default {
  name: 'Detail',
  props: {
    propBoardName: String
  },
  data () {
    return {
      boardName: this.propBoardName
    }
  },
  methods: {
    goBack () {
      if (this.boardName !== undefined) {
        this.$router.push(`/${this.boardName}`)
      }
    },
    isWriter () {
      const index = parseInt(this.$route.params.id)
      const post = this.$store.getters.getPostById(index)
      if (index === post.id) {
        return true
      }
      return false
    }
  },
  computed: {
    post () {
      const index = parseInt(this.$route.params.id)
      return this.$store.getters.getPostById(index)
    }
  }
}
</script>

<style scoped>

  .content-body{
    float: left;
    position:relative;
    border-right: 1px solid white !important;
    width: 700px;
    border-right: white;
    padding: 15px;
    box-sizing: border-box;
    display: block;
  }

  .content-function{
    float: right;
    position:relative;
    text-align: center;
    width: 100px;
    box-sizing: border-box;
    display: block;
  }

  .el-icon-delete{
    float: right;
    vertical-align: middle;
    font-size: 20px;
  }

  .name{
    font-size: 20px;
    font-weight: bold;
    margin-bottom: 5px;
    max-width: 200px;
    min-width: 50px;
  }

  .el-icon-back {
    cursor: pointer;
    font-size: 25px;
  }

  .content {
    border: 1px solid white;
    text-align: left;
    width: 800px;
    min-height: 800px;
    margin-top: 10px;
    margin-left: 10px;
  }

  .head{
    position: relative;
    display: block;
    box-sizing: border-box;
    padding: 10px;
    font-size: 10px;
    border-bottom: 1px solid white;
  }

  .content-main{
    position: relative;
    display: block;
    box-sizing: border-box;
  }

  .title{
    border-bottom: 1px solid white;
    font-size: 20px;
    font-weight: bold;
  }

  .text{
    width: 662px;
    overflow: auto;
    min-height: 180px;
    margin-top: 20px;
    font-size: 15px;

  }

</style>

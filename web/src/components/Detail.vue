<template>
  <div class="content clearfix">
    <div class="head clearfix">
      <i class="el-icon-back" @click="goBack"></i>
      <i class="el-icon-delete" @click="deletePost"></i>
      <i class="el-icon-edit" @click="editPost"></i>
      <div class="content-header">
        <div class="name">
          {{post.name}}
        </div>
        {{post.date}}
      </div>
    </div>
    <div class="content-main clearfix">
      <div class="content-body">
        <h2 class="title">
          {{post.title}}
        </h2>
        <hr>
        <div v-html="post.text" class="text">
        </div>
      </div>
      <div class="content-function">
        추천버튼 자리
      </div>
    </div>
  </div>
</template>

<script>
import _ from 'lodash'
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
    editPost () {
      this.$router.push('/edit/' + this.$route.params.id)
    },
    deletePost () {
      this.$confirm('정말로 게시글을 삭제하시겠습니까?', 'Warning', {
        confirmButtonText: '삭제',
        cancelButtonText: '취소',
        type: 'warning'
      }).then(() => {
        this.$message({
          type: 'success',
          message: '삭제 완료되었습니다.'
        })
        this.$http.delete(this.$http_url + '/api/post?id=' + this.$route.params.id)
          .then(() => {
            this.$router.push(`/${this.boardName}`)
          })
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '삭제 취소되었습니다.'
        })
      })
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
      const post = this.$store.getters.getPostById(index)
      const clonePost = _.cloneDeep(post)
      clonePost.text = clonePost.text.replace(/\n/gim, '<br />')
      return clonePost
    }
  }
}
</script>

<style scoped>

  .content-header{
    width: 200px;
  }

  .clearfix::after {
    content: "";
    clear: both;
    display: table;
  }

  .content-body{
    display: inline-block;
    border-right: 1px solid white !important;
    width: 700px;
    padding: 15px;
    box-sizing: border-box;
  }

  .content-function{
    display: inline-block;
    text-align: center;
    padding-top: 20px;
    width: 98px;
    box-sizing: border-box;
    vertical-align: top;
  }

  .el-icon-delete{
    clear: both;
    float: right;
    vertical-align: middle;
    font-size: 20px;
    cursor: pointer;
  }

  .el-icon-edit {
    clear: both;
    float: right;
    vertical-align: middle;
    font-size: 20px;
    cursor: pointer;
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
    width: 800px;
    margin-top: 10px;
    margin-left: 10px;
    box-sizing: border-box;
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
    overflow: hidden;
    box-sizing: border-box;
  }

  .title{
    font-size: 20px;
    font-weight: bold;
  }

  .text{
    min-height: 600px;
    width: 662px;
    margin-top: 20px;
    font-size: 15px;
  }

</style>

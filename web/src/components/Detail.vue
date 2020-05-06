<template>
  <el-card class="box-card">
    <div slot="header" class="clearfix">
      <el-row>
        <el-col :span="18">
          <div class="title"><i v-on:click="goBack" class="el-icon-back"></i> {{post.title}}</div>
        </el-col>
        <el-col :span="6"><div class="writerDate">{{post.name}}</div><div class="writerDate">{{post.date}}</div></el-col>
      </el-row>
    </div>
    <div class="textBox">
      <div class="text item">
        {{post.text}}
      </div>
    </div>
  </el-card>
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
        this.$router.push(`/main/${this.boardName}`)
      }
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
  .el-icon-back {
    cursor: pointer;
  }
  .text {
    font-size: 20px;
    text-align: left;
  }

  .item {
    margin-bottom: 18px;
  }

  .clearfix {
    text-align: left;
    color: white;
  }

  .box-card {
    margin-top: 30px;
    width: 1300px;
    height: 700px;
  }

  .title {
    margin-top: 30px;
    font-size: 30px;
    font-weight: bold;
    text-align: left;
  }

  .writerDate {
    margin-top: 20px;
    margin-bottom: 20px;
    font-size: 20px;
    text-align: right;
  }
</style>

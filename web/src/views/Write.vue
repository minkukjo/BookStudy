<template>
  <div class="write" style="margin-top: 30px">
    <el-form :model="ruleForm" status-icon :rules="rules" ref="ruleForm" label-width="120px" class="demo-ruleForm" style="width: 40%" >
      <el-form-item label="게시판 종류" prop="kind" >
        <el-select v-model="ruleForm.kind" placeholder="게시판을 선택해주세요." style="float:left;">
          <el-option label="자유 게시판" value="general"></el-option>
          <el-option label="스터디 모집" value="study"></el-option>
          <el-option label="질문 게시판" value="qna"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="제목" prop="title">
        <el-input type="title" v-model="ruleForm.title" autocomplete="off"></el-input>
      </el-form-item>
      <el-form-item label="내용" prop="text">
        <el-input type="textarea" v-model="ruleForm.text"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitForm('ruleForm')">작성</el-button>
        <el-button @click="resetForm('ruleForm')">초기화</el-button>
      </el-form-item>
    </el-form>
  </div>

</template>

<script>
export default {
  name: 'write',
  data () {
    return {
      ruleForm: {
        title: '',
        text: '',
        kind: ''
      },
      rules: {
        title: [
          { required: true, message: '제목을 입력해주세요.', trigger: 'blur' }
        ],
        text: [
          { required: true, message: '내용을 입력해주세요.', trigger: 'blur' }
        ],
        kind: [
          { required: true, message: '게시판을 선택해주세요.', trigger: 'change' }
        ]
      }
    }
  },
  methods: {
    submitForm (formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.$http.post(this.$http_url + '/api/write', this.ruleForm)
            .then(() => {
              if (this.ruleForm.kind === 'general') {
                this.$router.push('/main/general')
              } // else if (this.ruleForm.kind === 'study')
              // else if (this.ruleForm.kind === 'kind')
            })
          alert('작성 완료!')
        } else {
          console.log('error submit!!')
          return false
        }
      })
    },
    resetForm (formName) {
      this.$refs[formName].resetFields()
    }
  }

}
</script>

<style>
  .el-form-item label {
    color:white;
  }

  .el-textarea textarea {
    height: 300px;
  }
</style>

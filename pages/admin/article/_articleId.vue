<template>
  <div class="content-container">
    <client-only>
      <el-form ref="form" :model="form" :rules="rules" size="medium" label-width="100px">
        <el-form-item label="文章标题" prop="article_title">
          <el-input v-model="form.article_title" placeholder="请输入文章标题" :style="{width: '100%'}"/>
        </el-form-item>
        <el-form-item label="文章状态" prop="article_status">
          <el-select v-model="form.article_status" placeholder="请选择文章状态" :style="{width: '100%'}">
            <el-option label="发布" :value="1"/>
            <el-option label="草稿" :value="0"/>
          </el-select>
        </el-form-item>
        <el-form-item label="文章标签" prop="tag_ids">
          <el-select v-model="form.tag_ids" placeholder="请选择文章标签" multiple :style="{width: '100%'}"
                     @change="$forceUpdate()">
            <el-option
              v-for="(item, index) in tags"
              :key="index"
              :label="item.tag_name"
              :value="item.tag_id"
            >
              <span style="float: left">{{ item.tag_name }}</span>
              <img :src="item.tag_icon" style="float: right;width: 27px;height: 25px">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="文章描述" prop="article_desc">
          <el-input
            v-model="form.article_desc"
            type="textarea"
            placeholder="请输入文章描述"
            resize="none"
            rows="4"
            maxlength="300"
            :style="{width: '100%'}"
          />
        </el-form-item>
        <el-form-item label="文章内容" prop="article_content">
          <client-only>
            <tinymce v-model="form.article_content" placeholder="请输入文章内容" :height="500"/>
          </client-only>
        </el-form-item>
        <el-form-item size="large">
          <el-button type="primary" @click="submitForm">
            保 存
          </el-button>
          <el-button @click="reset">
            重 置
          </el-button>
        </el-form-item>
      </el-form>
    </client-only>
  </div>
</template>

<script>
  import {Message} from 'element-ui'
  import Tinymce from '~/components/tinymce/index.vue'
  import {tagList, updateArticle, addArticle, getArticle} from '~/apis/admin'

  export default {
    name: 'AdminArticleDetail',
    components: {Tinymce},
    layout: 'admin',
    data() {
      return {
        articleId: 0,
        tags: [],
        form: {
          article_title: undefined,
          tag_ids: [],
          article_status: 1,
          article_desc: undefined,
          article_content: ''
        },
        rules: {
          article_title: [{
            required: true,
            message: '请输入文章标题',
            trigger: 'blur'
          }],
          tag_ids: [{
            required: true,
            type: 'array',
            message: '请至少选择一个标签',
            trigger: 'change'
          }],
          article_content: [{
            required: true,
            message: '请输入文章内容',
            trigger: 'blur'
          }]
        }
      }
    },
    head() {
      return {
        title: this.articleId > 0 ? '修改文章 - 后台管理' : '添加文章 - 后台管理'
      }
    },
    created() {
      this.articleId = parseInt(this.$route.params.articleId)
    },
    mounted() {
      if (this.articleId > 0) {
        getArticle(this.articleId).then((response) => {
          if (response.code !== 200) {
            return Message.error(response.message)
          }
          this.form = {
            article_title: response.data.article_title,
            article_status: response.data.article_status,
            article_desc: response.data.article_desc,
            article_content: response.data.article_content,
            tag_ids: response.data.tags.map((item) => {
              return item.tag_id
            })
          }
        })
      }
      tagList().then((response) => {
        if (response.code !== 200) {
          return Message.error(response.message)
        }
        this.tags = response.data
      })
    },
    methods: {
      reset() {
        this.form = {
          article_title: undefined,
          tag_ids: [],
          article_status: 1,
          article_desc: undefined,
          article_content: ''
        }
      },
      submitForm() {
        this.$refs.form.validate((valid) => {
          if (valid) {
            if (this.articleId > 0) {
              updateArticle(this.articleId, this.form).then((response) => {
                if (response.code !== 200) {
                  return Message.error(response.message)
                }
                Message.success(response.message)
              })
            } else {
              addArticle(this.form).then((response) => {
                if (response.code !== 200) {
                  return Message.error(response.message)
                }
                Message.success(response.message)
                setTimeout(() => {
                  this.$router.push({path: '/admin/article'})
                }, 1500)
              })
            }
          }
        })
      }
    }
  }
</script>

<style scoped>
  button {
    width: 100px;
  }

  /* ::v-deep .tox {
    color: var(--color);
  }
  ::v-deep .tox-tinymce,
  ::v-deep .tox:not([dir=rtl]) .tox-toolbar__group:not(:last-of-type) {
    border-color: var(--border-color);
  }

  ::v-deep .tox .tox-toolbar,
  ::v-deep .tox .tox-toolbar__overflow,
  ::v-deep .tox .tox-toolbar__primary,
  ::v-deep .tox .tox-edit-area__iframe {
    background-color: var(--bg);
  }

  ::v-deep .tox .tox-tbtn:hover {
    background-color: var(--bg-secondary);
  } */

</style>

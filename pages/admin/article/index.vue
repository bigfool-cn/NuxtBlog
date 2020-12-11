<template>
  <div class="content-container">
    <client-only>
      <div class="serach-container">
        <el-input v-model="query.article_title" style="width: 300px" placeholder="文章标题" />
        <el-select v-model="query.article_status" placeholder="文章状态">
          <el-option label="全部" :value="-1" />
          <el-option label="发布" :value="1" />
          <el-option label="草稿" :value="0" />
        </el-select>
        <el-select v-model="query.tag_id" placeholder="文章标签" style="width: 200px">
          <el-option label="全部" :value="-1" />
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
        <el-button @click="handleSearch">
          搜索
        </el-button>
      </div>
      <div class="table-container">
        <div class="table-btn">
          <el-button
            type="danger"
            icon="el-icon-delete"
            size="mini"
            :disabled="multiple"
            @click="handleDelete"
          >
            删除
          </el-button>
          <el-button
            type="primary"
            icon="el-icon-plus"
            size="mini"
            @click="handleAdd"
          >
            新增
          </el-button>
        </div>

        <el-table :data="articleList" border @selection-change="handleSelectionChange">
          <el-table-column type="selection" width="55" align="center" />
          <el-table-column label="名称" prop="article_title" :show-overflow-tooltip="true" />
          <el-table-column label="标签" prop="tag_icon">
            <template slot-scope="scope">
              <span v-for="(tag,key) in scope.row.tags" :key="key" :title="tag.tag_name">
                <img :src="tag.tag_icon" style="width:27px; height:25px" >
              </span>
            </template>
          </el-table-column>
          <el-table-column label="阅读量" prop="article_read" width="80" />
          <el-table-column label="修改时间" align="center" prop="update_time">
            <template slot-scope="scope">
              <span>{{ scope.row.update_time | parseTime }}</span>
            </template>
          </el-table-column>
          <el-table-column label="创建时间" align="center" prop="create_time">
            <template slot-scope="scope">
              <span>{{ scope.row.create_time | parseTime }}</span>
            </template>
          </el-table-column>
          <el-table-column label="操作" align="center" class-name="small-padding fixed-width" width="180">
            <template slot-scope="scope">
              <el-button
                size="mini"
                type="text"
                icon="el-icon-edit"
                @click="handleUpdate(scope.row)"
              >
                修改
              </el-button>
              <el-button
                size="mini"
                type="text"
                icon="el-icon-delete"
                @click="handleDelete(scope.row)"
              >
                删除
              </el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
      <div class="pagination" v-if="total / query.limit > 1">
        <el-pagination
          background
          layout="prev, pager, next"
          :page-size="query.limit"
          :total="total">
        </el-pagination>
      </div>
    </client-only>
  </div>
</template>

<script>
import { articleList, tagList } from '~/apis/admin'

export default{
  name: 'AdminArticle',
  layout: 'admin',
  data () {
    return {
      multiple: true,
      selectIds: [],
      total: 0,
      query: {
        page: 0,
        limit: 20,
        article_title: '',
        article_status: -1,
        tag_id: -1
      },
      tags: [],
      articleList: []
    }
  },
  head () {
    return {
      title: '文章 - 后台管理'
    }
  },
  mounted () {
    this.getArticleList()
    this.getTagList()
  },
  methods: {
    getArticleList () {
      articleList(this.query).then((response) => {
        if (response.code !== 200) {
          return this.$toast.error(response.message)
        }
        this.articleList = response.data.articles
        this.total = response.data.total
      })
    },
    handleSearch () {
      this.getArticleList()
    },
    getTagList () {
      tagList(this.query).then((response) => {
        if (response.code !== 200) {
          return this.$toast.error(response.message)
        }
        this.tags = response.data
      })
    },
    handleSelectionChange (rows) {
      this.selectIds = rows.map(item => item.article_id)
      this.multiple = false
    },
    handleAdd () {
      this.$router.push({ path: '/admin/article/add' })
    },
    handleUpdate (row) {
      this.$router.push({ path: '/admin/article/' + row.article_id })
    },
    handleDelete (row) {
      console.log(row)
    }
  }
}
</script>

<style scoped>
  .pagination {
    margin-top: 15px;
    text-align: center;
  }
  button {
    width: 100px;
  }
</style>

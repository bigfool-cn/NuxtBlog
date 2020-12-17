<template>
  <div class="content-container">
    <client-only>
      <div class="serach-container">
        <el-input v-model="query.tag_name" style="width: 300px" placeholder="标签名称" />
        <el-button @click="handleSearch">搜索</el-button>
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

        <el-table :data="tagList" border @selection-change="handleSelectionChange">
          <el-table-column type="selection" width="55" align="center" />
          <el-table-column label="图标" prop="tag_icon" width="100">
            <template slot-scope="scope">
              <img :src="scope.row.tag_icon">
            </template>
          </el-table-column>
          <el-table-column label="名称" prop="tag_name" :show-overflow-tooltip="true" />
          <el-table-column label="创建时间" align="center" prop="create_time">
            <template slot-scope="scope">
              <span>{{ scope.row.create_time | parseTime }}</span>
            </template>
          </el-table-column>
          <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
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
      <el-dialog
        v-drag-dialog
        :title="title"
        :visible.sync="dialogVisible"
        center
        width="30%"
      >
        <el-form ref="form" :model="form" :rules="rules" label-width="80px">
          <el-form-item label="图标" prop="tag_icon">
            <input id="file" type="file" name="file" style="display:none" @change="changeFile">
            <img :src="form.tag_icon ? form.tag_icon : icon" class="tag-icon" @click="uploadIcon">
            <el-input v-model="form.tag_icon" type="hidden" />
          </el-form-item>
          <el-form-item label="名称" prop="tag_name" class="tag-name">
            <el-input v-model="form.tag_name" placeholder="名称" maxlength="15" />
          </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button type="primary" @click="submitForm">
            确 定
          </el-button>
          <el-button>
            取 消
          </el-button>
        </div>
      </el-dialog>
    </client-only>
  </div>
</template>

<script>
import { Message } from 'element-ui'
import { tagList, updateTag, addTag, delTag } from '~/apis/admin'
export default {
  name: 'AdminTag',
  layout: 'admin',
  data () {
    return {
      dialogVisible: false,
      multiple: true,
      title: '',
      selectIds: [],
      query: {
        page: 0,
        limit: 20,
        tag_name: ''
      },
      icon: require('~/assets/img/upload.png'),
      form: {
        tag_id: 0,
        tag_icon: '',
        tag_name: ''
      },
      rules: {
        tag_name: [
          { required: true, message: '名称不能为空', trigger: 'blur' }
        ]
      },
      tagList: []
    }
  },
  head () {
    return {
      title: '标签 - 后台管理'
    }
  },
  mounted () {
    this.getTagList()
  },
  methods: {
    getTagList () {
      tagList(this.query).then((response) => {
        if (response.code !== 200) {
          return Message.error(response.message)
        }
        this.tagList = response.data
      })
    },
    handleSearch () {
      this.getTagList()
    },
    handleSelectionChange (rows) {
      this.selectIds = rows.map(item => item.tag_id)
      this.multiple = false
    },
    handleAdd () {
      this.title = '新增标签'
      this.dialogVisible = true
      this.reset()
    },
    handleUpdate (row) {
      this.reset()
      this.title = '修改标签'
      this.dialogVisible = true
      this.form = {
        tag_id: row.tag_id,
        tag_icon: row.tag_icon,
        tag_name: row.tag_name
      }
      this.icon = row.tag_icon
    },
    handleDelete (row) {
      const ids = row.tag_id ? [row.tag_id] : this.selectIds
      delTag(ids).then((response) => {
        if (response.code !== 200) {
          return Message.error(response.message)
        }
        Message.success(response.message)
        this.getTagList()
      })
    },
    reset () {
      this.form = {
        tag_id: 0,
        tag_icon: '',
        tag_name: ''
      }
      this.icon = require('~/assets/img/upload.png')
    },
    uploadIcon () {
      const fileObj = document.querySelector('#file')
      fileObj.click()
    },
    changeFile (event) {
      const file = event.target.files.item(0)

      if (!file) {
        return Message.error('未选择文件')
      }

      if (file.size > 1024 * 1024) {
        return Message.error('图标不能超多1MB')
      }

      const fr = new FileReader()
      fr.onload = (e) => {
        this.form.tag_icon = e.target.result
        const obj = document.querySelector('.tag-icon')
        obj.src = this.form.tag_icon
      }
      fr.readAsDataURL(file)

      event.target.value = ''
    },
    submitForm () {
      this.$refs.form.validate((valid) => {
        if (!this.form.tag_icon) {
          return Message.error('请上传图标')
        }
        if (valid) {
          if (this.form.tag_id) {
            updateTag(this.form.tag_id, this.form).then((response) => {
              if (response.code !== 200) {
                return Message.error(response.message)
              }
              Message.success(response.message)
              this.dialogVisible = false
              this.getTagList()
            })
          } else {
            addTag(this.form).then((response) => {
              if (response.code !== 200) {
                return Message.error(response.message)
              }
              Message.success(response.message)
              this.dialogVisible = false
              this.getTagList()
            })
          }
        }
      })
    }
  }
}
</script>

<style scoped>
  .tag-icon {
    position: absolute;
    cursor: pointer;
  }
  .tag-name {
    padding-top: 10px;
  }
  button {
    width: 100px;
  }
</style>

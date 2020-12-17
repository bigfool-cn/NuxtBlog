<template>
  <div class="content-container">
    <client-only>
      <div class="serach-container">
        <el-input
          v-model="query.vistor_ip"
          style="width: 300px"
          placeholder="IP"
        />
        <el-date-picker
          v-model="date"
          :default-time="['00:00:00', '23:59:59']"
          type="datetimerange"
          range-separator="至"
          start-placeholder="开始日期"
          end-placeholder="结束日期"
          format="yyyy-MM-dd HH:mm:ss"
          value-format="yyyy-MM-dd HH:mm:ss"
        >
        </el-date-picker>
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
        </div>

        <el-table
          :data="vistorList"
          border
          @selection-change="handleSelectionChange"
        >
          <el-table-column type="selection" width="55" align="center" />
          <el-table-column
            label="IP"
            prop="vistor_ip"
            :show-overflow-tooltip="true"
          />
          <el-table-column
            label="Path"
            prop="vistor_path"
            :show-overflow-tooltip="true"
          />
          <el-table-column
            label="UserAgent"
            prop="user_agent"
            :show-overflow-tooltip="true"
          />
          <el-table-column label="创建时间" align="center" prop="create_time">
            <template slot-scope="scope">
              <span>{{ scope.row.create_time | parseTime }}</span>
            </template>
          </el-table-column>
          <el-table-column
            label="操作"
            align="center"
            class-name="small-padding fixed-width"
          >
            <template slot-scope="scope">
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
          layout="total, prev, pager, next"
          :page-size="query.limit"
          :total="total"
          @current-change="toPage"
        >
        </el-pagination>
      </div>
    </client-only>
  </div>
</template>

<script>
import { Message } from "element-ui";
import { vistorList, delVistor } from "~/apis/admin";
export default {
  layout: "admin",
  data() {
    return {
      multiple: true,
      selectIds: [],
      query: {
        page: 0,
        limit: 20,
        vistor_ip: "",
        start_time: "",
        end_time: ""
      },
      date:undefined,
      total: 0,
      vistorList: [],
    };
  },
  head() {
    return {
      title: "后台管理",
    };
  },
  watch: {
    date(val) {
      if(val !== null) {
        console.log(val)
        this.query.start_time = val[0]
        this.query.end_time = val[1]
      } else {
        this.query.start_time = ""
        this.query.end_time = ""
      }
    }
  },
  mounted() {
    this.getVistorList();
  },
  methods: {
    getVistorList() {
      vistorList(this.query).then((response) => {
        if (response.code !== 200) {
          return Message.error(response.message);
        }
        this.vistorList = response.data.vistors;
        this.total = response.data.total;
      });
    },
    handleSearch() {
      this.getVistorList();
    },
    handleSelectionChange(rows) {
      this.selectIds = rows.map((item) => item.vistor_id);
      this.multiple = false;
    },
    handleDelete(row) {
      const ids = row.vistor_id ? [row.vistor_id] : this.selectIds;
      delVistor(ids).then((response) => {
        if (response.code !== 200) {
          return Message.error(response.message);
        }
        Message.success(response.message);
        this.getVistorList();
      });
    },
    toPage(page) {
      this.query.page = page;
      this.getVistorList();
    },
  },
};
</script>

<style scoped>
.pagination {
  margin-top: 15px;
  text-align: center;
}

button {
  width: 100px;
}

::v-deep .el-date-editor .el-range-input,
::v-deep .el-picker-panel {
  background-color: var(--bg);
  color: var(--color);
}

::v-deep .el-pagination.is-background .btn-prev,
::v-deep .el-pagination.is-background .btn-next,
::v-deep .el-pagination.is-background .el-pager li {
  color: var(--color);
  background-color: var(--bg-secondary);
}

::v-deep .el-pagination.is-background .el-pager li:not(.disabled).active {
  color: var(--color-active);
  background-color: var(--bg-secondary);
}
</style>

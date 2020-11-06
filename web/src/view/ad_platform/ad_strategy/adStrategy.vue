<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="用户类型">
          <el-input placeholder="搜索条件" v-model="searchInfo.userRole"></el-input>
        </el-form-item>
        <el-form-item label="平台">
          <el-input placeholder="搜索条件" v-model="searchInfo.platform"></el-input>
        </el-form-item>
        <el-form-item label="生效 " prop="status">
          <el-select v-model="searchInfo.status" clear placeholder="请选择">
            <el-option
                key="true"
                label="是"
                value="true">
            </el-option>
            <el-option
                key="false"
                label="否"
                value="false">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="操作人">
          <el-input placeholder="搜索条件" v-model="searchInfo.operator"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增adStrategy表</el-button>
        </el-form-item>
        <el-form-item>
          <el-popover placement="top" v-model="deleteVisible" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button @click="deleteVisible = false" size="mini" type="text">取消</el-button>
              <el-button @click="onDelete" size="mini" type="primary">确定</el-button>
            </div>
            <el-button icon="el-icon-delete" size="mini" slot="reference" type="danger">批量删除</el-button>
          </el-popover>
        </el-form-item>
      </el-form>
    </div>
    <el-table
        :data="tableData"
        @selection-change="handleSelectionChange"
        border
        ref="multipleTable"
        stripe
        style="width: 100%"
        tooltip-effect="dark"
    >
      <el-table-column type="selection" width="55"></el-table-column>

      <el-table-column label="用户类型" prop="userRole" width="120"></el-table-column>

      <el-table-column label="平台" prop="platform" width="120"></el-table-column>

      <el-table-column label="生效 " prop="status" width="120">
        <template slot-scope="scope">{{ scope.row.status|formatBoolean }}</template>
      </el-table-column>

      <el-table-column label="广告策略" prop="strategy" width="420">
        <template slot-scope="scope">
          <json-viewer
              :value="parseJsonBody(scope.row.strategy)"
              :expand-depth=5
              copyable
              boxed
              sort></json-viewer>
        </template>
      </el-table-column>

      <el-table-column label="创建时间" prop="createTime" width="120"></el-table-column>

      <el-table-column label="更新时间" prop="updateTime" width="120"></el-table-column>

      <el-table-column label="操作人" prop="operator" width="120"></el-table-column>

      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateAdStrategy(scope.row)" size="small" type="primary"
                     icon="el-icon-edit">变更
          </el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteAdStrategy(scope.row)">确定</el-button>
            </div>
            <el-button type="danger" icon="el-icon-delete" size="mini" slot="reference">删除</el-button>
          </el-popover>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
        :current-page="page"
        :page-size="pageSize"
        :page-sizes="[10, 30, 50, 100]"
        :style="{float:'right',padding:'20px'}"
        :total="total"
        @current-change="handleCurrentChange"
        @size-change="handleSizeChange"
        layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="用户类型:">
          <el-input v-model="formData.userRole" clearable placeholder="请输入"></el-input>
        </el-form-item>

        <el-form-item label="平台:">
          <el-input v-model="formData.platform" clearable placeholder="请输入"></el-input>
        </el-form-item>

        <el-form-item label="生效 :">
          <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否"
                     v-model="formData.status" clearable></el-switch>
        </el-form-item>

        <el-form-item label="广告策略:">
          <JsonEditor
              :options="{
            confirmText: 'confirm',
            cancelText: 'cancel',
        }"
              :objData="parseJsonBody(formData.strategy)"
              v-model="formData.strategy">
          </JsonEditor>
        </el-form-item>

        <el-form-item label="创建时间:">
          <el-date-picker type="date" placeholder="选择日期" v-model="formData.createTime" clearable></el-date-picker>
        </el-form-item>

        <el-form-item label="更新时间:">
          <el-date-picker type="date" placeholder="选择日期" v-model="formData.updateTime" clearable></el-date-picker>
        </el-form-item>

        <el-form-item label="操作人:">
          <el-input v-model="formData.operator" clearable placeholder="请输入"></el-input>
        </el-form-item>
      </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
  createAdStrategy,
  deleteAdStrategy,
  deleteAdStrategyByIds,
  updateAdStrategy,
  findAdStrategy,
  getAdStrategyList
} from "@/api/adStrategy";  //  此处请自行替换地址
import {formatTimeToStr} from "@/utils/date";
import infoList from "@/mixins/infoList";

export default {
  name: "AdStrategy",
  mixins: [infoList],
  data() {
    return {
      listApi: getAdStrategyList,
      dialogFormVisible: false,
      visible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [], formData: {
        userRole: "",
        platform: "",
        status: false,
        strategy: "",
        createTime: new Date(),
        updateTime: new Date(),
        operator: "",
      },
      jsonData: {
        name: 'mike',
        age: 23,
        phone: '18552129932',
        address: ['AAA C1', 'BBB C2']
      }
    };
  },
  filters: {
    formatDate: function (time) {
      if (time != null && time != "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
      } else {
        return "";
      }
    },
    formatBoolean: function (bool) {
      if (bool != null) {
        return bool ? "是" : "否";
      } else {
        return "";
      }
    }
  },
  methods: {
    //条件搜索前端看此方法
    onSubmit() {
      this.page = 1
      this.pageSize = 10
      if (this.searchInfo.status == "") {
        this.searchInfo.status = null
      }
      this.getTableData()
    },
    handleSelectionChange(val) {
      this.multipleSelection = val
    },
    async onDelete() {
      const ids = []
      this.multipleSelection &&
      this.multipleSelection.map(item => {
        ids.push(item.ID)
      })
      const res = await deleteAdStrategyByIds({ids})
      if (res.code == 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        this.deleteVisible = false
        this.getTableData()
      }
    },
    async updateAdStrategy(row) {
      const res = await findAdStrategy({ID: row.ID});
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.readStrategy;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
        userRole: "",
        platform: "",
        status: false,
        strategy: "",
        createTime: new Date(),
        updateTime: new Date(),
        operator: "",

      };
    },
    async deleteAdStrategy(row) {
      this.visible = false;
      const res = await deleteAdStrategy({ID: row.ID});
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功"
        });
        this.getTableData();
      }
    },
    async enterDialog() {
      let res;
      this.formData.strategy = JSON.stringify(this.formData.strategy)
      switch (this.type) {
        case "create":
          res = await createAdStrategy(this.formData);
          break;
        case "update":
          res = await updateAdStrategy(this.formData);
          break;
        default:
          res = await createAdStrategy(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "创建/更改成功"
        })
        this.closeDialog();
        this.getTableData();
      }
    },
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    },
    parseJsonBody(value) {
      try {
        return JSON.parse(value)
      } catch (err) {
        return value
      }
    }
  },
  async created() {
    await this.getTableData();

  }
};
</script>

<style>
</style>
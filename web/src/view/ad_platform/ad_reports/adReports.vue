<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter.native="onSubmit"
               @keyup.esc.native="resetDialog">
        <el-form-item label="平台">
          <el-input placeholder="搜索条件" v-model="searchInfo.platform"></el-input>
        </el-form-item>
        <el-form-item label="广告id">
          <el-input placeholder="搜索条件" v-model="searchInfo.adId"></el-input>
        </el-form-item>
        <el-form-item label="广告类型">
          <el-input placeholder="搜索条件" v-model="searchInfo.adType"></el-input>
        </el-form-item>
        <el-form-item label="用户id">
          <el-input placeholder="搜索条件" v-model="searchInfo.userId"></el-input>
        </el-form-item>
        <el-form-item label="设备id">
          <el-input placeholder="搜索条件" v-model="searchInfo.guestId"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="resetDialog" type="primary">重置</el-button>
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
        @sort-change="onOrder"
    >
      <el-table-column type="selection" width="55"></el-table-column>

      <el-table-column label="平台" prop="platform" width="120" sortable></el-table-column>

      <el-table-column label="广告id" prop="adId" width="240" sortable></el-table-column>

      <el-table-column label="广告类型" prop="adType" width="120" v-model="adTypeMap" sortable>
        <template slot-scope="scope">{{ adTypeMap[scope.row.adType] }}</template>
      </el-table-column>

      <el-table-column label="用户id" prop="userId" width="120"></el-table-column>

      <el-table-column label="设备id" prop="guestId" width="240"></el-table-column>

      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateAdReports(scope.row)" size="small" type="primary"
                     icon="el-icon-edit">变更
          </el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteAdReports(scope.row)">确定</el-button>
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
        <el-form-item label="平台:">
          <el-input v-model="formData.platform" clearable placeholder="请输入"></el-input>
        </el-form-item>

        <el-form-item label="广告id:">
          <el-input v-model="formData.adId" clearable placeholder="请输入"></el-input>
        </el-form-item>

        <!--        <el-form-item label="广告类型:">-->
        <!--          <el-input v-model="formData.adType" clearable placeholder="请输入"></el-input>-->
        <!--        </el-form-item>-->
        <el-form-item label="广告类型:" prop="adType">
          <el-select v-model="formData.adType" placeholder="请输入" clearable :style="{width: '100%'}">
            <el-option v-for="(item, index) in adTypeOptions" :key="index" :label="item.label"
                       :value="item.value" :disabled="item.disabled"></el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="用户id:">
          <el-input v-model.number="formData.userId" clearable placeholder="请输入"></el-input>
        </el-form-item>

        <el-form-item label="设备id:">
          <el-input v-model="formData.guestId" clearable placeholder="请输入"></el-input>
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
  createAdReports,
  deleteAdReports,
  deleteAdReportsByIds,
  updateAdReports,
  findAdReports,
  getAdReportsList
} from "@/api/adReports";  //  此处请自行替换地址
import {formatTimeToStr} from "@/utils/date";
import infoList from "@/mixins/infoList";

export default {
  name: "AdReports",
  mixins: [infoList],
  data() {
    return {
      listApi: getAdReportsList,
      dialogFormVisible: false,
      visible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [], formData: {
        platform: "",
        adId: "",
        adType: "",
        userId: 0,
        guestId: "",
      },
      adTypeOptions: [{
        "label": "视频",
        "value": "video"
      }, {
        "label": "横幅",
        "value": "banner"
      }, {
        "label": "激励广告",
        "value": "gift"
      }],
      adTypeMap: {
        "video": "视频",
        "banner": "横幅",
        "gift": "激励广告",
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
      this.getTableData()
    },
    onOrder(scope) {
      this.order = scope.prop;
      this.orderBy = scope.order === "descending" ? "desc" : "asc";
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
      const res = await deleteAdReportsByIds({ids})
      if (res.code == 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        this.deleteVisible = false
        this.getTableData()
      }
    },
    async updateAdReports(row) {
      const res = await findAdReports({ID: row.ID});
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.readReports;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
        platform: "",
        adId: "",
        adType: "",
        userId: 0,
        guestId: "",

      };
    },
    async deleteAdReports(row) {
      this.visible = false;
      const res = await deleteAdReports({ID: row.ID});
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
      switch (this.type) {
        case "create":
          res = await createAdReports(this.formData);
          break;
        case "update":
          res = await updateAdReports(this.formData);
          break;
        default:
          res = await createAdReports(this.formData);
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
    resetDialog() {
      this.searchInfo = {};
    }
  },
  async created() {
    await this.getTableData();

  }
};
</script>

<style>
</style>
<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-select v-model="query.categoryID" :placeholder="t('view.kpi.item.category')" style="width: 220px" @change="fetchTable">
          <el-option v-for="c in categories" :key="c.ID" :label="c.name" :value="c.ID" />
        </el-select>
        <el-button type="primary" icon="plus" @click="openDrawer">{{ t('general.add') }}</el-button>
        <el-popconfirm :title="t('general.deleteConfirm')" @confirm="onBatchDelete">
          <template #reference>
            <el-button icon="delete">{{ t('general.delete') }}</el-button>
          </template>
        </el-popconfirm>
      </div>
      <el-table
        ref="multipleTable"
        :data="tableData"
        style="width: 100%"
        tooltip-effect="dark"
        row-key="ID"
        @selection-change="onSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" :label="t('general.createdAt')" width="180">
          <template #default="scope">
            <span>{{ formatDate(scope.row.CreatedAt) }}</span>
          </template>
        </el-table-column>
        <el-table-column align="left" :label="t('view.kpi.item.name')" prop="name" min-width="160" />
        <el-table-column align="left" :label="t('view.kpi.item.weight')" prop="weight" width="120" />
        <el-table-column align="left" :label="t('general.operations')" min-width="160">
          <template #default="scope">
            <el-button link type="primary" icon="edit" @click="onEdit(scope.row)">{{ t('general.edit') }}</el-button>
            <el-popconfirm :title="t('general.deleteConfirm')" @confirm="onDelete(scope.row)">
              <template #reference>
                <el-button link type="danger" icon="delete">{{ t('general.delete') }}</el-button>
              </template>
            </el-popconfirm>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 20, 50, 100]"
          layout="total, sizes, prev, pager, next, jumper"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>

    <el-drawer v-model="drawerVisible" :before-close="closeDrawer" :show-close="false">
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{ t('view.kpi.item.title') }}</span>
          <div>
            <el-button @click="closeDrawer">{{ t('general.cancel') }}</el-button>
            <el-button type="primary" @click="submitForm">{{ t('general.confirm') }}</el-button>
          </div>
        </div>
      </template>

      <el-form ref="formRef" :model="form" :rules="rules" label-width="160px">
        <el-form-item :label="t('view.kpi.item.category')" prop="categoryID">
          <el-select v-model="form.categoryID" placeholder="Category" style="width: 100%">
            <el-option v-for="c in categories" :key="c.ID" :label="c.name" :value="c.ID" />
          </el-select>
        </el-form-item>
        <el-form-item :label="t('view.kpi.item.name')" prop="name">
          <el-input v-model="form.name" />
        </el-form-item>
        <el-form-item :label="t('view.kpi.item.weight')" prop="weight">
          <el-input-number v-model="form.weight" :min="0" :max="100" />
        </el-form-item>
      </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { ElMessage } from 'element-plus'
import { formatDate } from '@/utils/format'
import { getKpiCategoryList } from '@/api/kpiCategory'
import { 
  createKpiItem, updateKpiItem, deleteKpiItem, deleteKpiItemByIds, findKpiItem, getKpiItemList
} from '@/api/kpiItem'

defineOptions({ name: 'KpiItem' })

const { t } = useI18n()
const formRef = ref(null)

const categories = ref([])
const query = ref({ categoryID: undefined })

const form = ref({ ID: undefined, categoryID: undefined, name: '', weight: 0 })

const rules = {
  categoryID: [{ required: true, message: t('general.cannotBeEmpty'), trigger: 'change' }],
  name: [{ required: true, message: t('general.cannotBeEmpty'), trigger: 'blur' }],
  weight: [{ required: true, message: t('general.cannotBeEmpty'), trigger: 'change' }]
}

const tableData = ref([])
const total = ref(0)
const page = ref(1)
const pageSize = ref(10)
const multipleSelection = ref([])

const handleSizeChange = (val) => { pageSize.value = val; fetchTable() }
const handleCurrentChange = (val) => { page.value = val; fetchTable() }

const loadCategories = async () => {
  const res = await getKpiCategoryList({ page: 1, pageSize: 1000 })
  if (res.code === 0) categories.value = res.data.list
}
loadCategories()

const fetchTable = async () => {
  const res = await getKpiItemList({ page: page.value, pageSize: pageSize.value, categoryID: query.value.categoryID })
  if (res.code === 0) {
    tableData.value = res.data.list
    total.value = res.data.total
    page.value = res.data.page
    pageSize.value = res.data.pageSize
  }
}
fetchTable()

const drawerVisible = ref(false)
const type = ref('create')

const resetForm = () => { form.value = { ID: undefined, categoryID: query.value.categoryID, name: '', weight: 0 } }
const openDrawer = () => { type.value = 'create'; resetForm(); drawerVisible.value = true }
const closeDrawer = () => { drawerVisible.value = false; resetForm() }

const onEdit = async (row) => {
  const res = await findKpiItem({ ID: row.ID })
  if (res.code === 0) {
    const d = res.data.reItem
    form.value = { ID: d.ID, categoryID: d.categoryID, name: d.name, weight: d.weight }
    type.value = 'update'
    drawerVisible.value = true
  }
}

const submitForm = () => {
  formRef.value.validate(async (valid) => {
    if (!valid) return
    let res
    if (type.value === 'create') res = await createKpiItem(form.value)
    else res = await updateKpiItem(form.value)
    if (res.code === 0) { ElMessage.success(t('general.createUpdateSuccess')); closeDrawer(); fetchTable() }
  })
}

const onDelete = (row) => async () => {
  const res = await deleteKpiItem({ ID: row.ID })
  if (res.code === 0) { ElMessage.success(t('general.deleteSuccess')); if (tableData.value.length === 1 && page.value > 1) page.value--; fetchTable() }
}

const onSelectionChange = (rows) => { multipleSelection.value = rows }
const onBatchDelete = async () => {
  if (!multipleSelection.value.length) return
  const ids = multipleSelection.value.map(i => i.ID)
  const res = await deleteKpiItemByIds({ ids })
  if (res.code === 0) { ElMessage.success(t('general.deleteSuccess')); if (tableData.value.length === ids.length && page.value > 1) page.value--; fetchTable() }
}
</script>



<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
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
        <el-table-column align="left" :label="t('view.rating.rating')" prop="rating" width="120">
          <template #default="scope">
            <el-rate :model-value="scope.row.rating" :max="5" disabled/>
          </template>
        </el-table-column>
        <el-table-column align="left" :label="t('view.rating.description')" prop="description" width="200" />
        <el-table-column align="left" :label="t('view.rating.performanceLevel')" prop="performanceLevel" width="240" />
        <el-table-column align="left" :label="t('view.rating.range')" min-width="180">
          <template #default="scope">
            <span>{{ renderRange(scope.row.minPercentage, scope.row.maxPercentage) }}</span>
          </template>
        </el-table-column>
        <el-table-column align="left" :label="t('general.operations')" min-width="160">
          <template #default="scope">
            <el-button type="primary" link icon="edit" @click="onEdit(scope.row)">{{ t('general.change') }}</el-button>
            <el-popconfirm :title="t('general.deleteConfirm')" @confirm="onDelete(scope.row)">
              <template #reference>
                <el-button type="primary" link icon="delete">{{ t('general.delete') }}</el-button>
              </template>
            </el-popconfirm>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>

    <el-drawer v-model="drawerVisible" :before-close="closeDrawer" :show-close="false">
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{ t('view.rating.title') }}</span>
          <div>
            <el-button @click="closeDrawer">{{ t('general.cancel') }}</el-button>
            <el-button type="primary" @click="submitForm">{{ t('general.confirm') }}</el-button>
          </div>
        </div>
      </template>

      <el-form ref="formRef" :model="form" :rules="rules" label-width="160px">
        <el-form-item :label="t('view.rating.rating')" prop="rating">
          <el-rate v-model="form.rating" :max="5" />
        </el-form-item>

        <el-form-item :label="t('view.rating.description')" prop="description">
          <el-input v-model="form.description" />
        </el-form-item>

        <el-form-item :label="t('view.rating.performanceLevel')" prop="performanceLevel">
          <el-input v-model="form.performanceLevel" />
        </el-form-item>

        <el-form-item :label="t('view.rating.minPercentage')" prop="minPercentage">
          <el-input-number v-model="form.minPercentage" :min="0" :max="100" />
        </el-form-item>

        <el-form-item :label="t('view.rating.maxPercentage')" prop="maxPercentage">
          <el-input-number v-model="form.maxPercentage" :min="0" :max="100" />
        </el-form-item>
      </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { ElMessage, ElMessageBox } from 'element-plus'
import { formatDate } from '@/utils/format'
import {
  createRating,
  updateRating,
  deleteRating,
  deleteRatingByIds,
  findRating,
  getRatingList
} from '@/api/rating'

const { t } = useI18n()

defineOptions({ name: 'Rating' })

const formRef = ref(null)

const form = ref({
  ID: undefined,
  rating: null,
  description: '',
  performanceLevel: '',
  minPercentage: 0,
  maxPercentage: 0
})

const rules = {
  rating: [
    { required: true, message: t('general.cannotBeEmpty'), trigger: 'change' },
    { type: 'number', min: 1, max: 5, message: t('general.hint') }
  ],
  description: [{ required: true, message: t('general.cannotBeEmpty'), trigger: 'blur' }],
  performanceLevel: [{ required: true, message: t('general.cannotBeEmpty'), trigger: 'blur' }],
  minPercentage: [
    { required: true, message: t('general.cannotBeEmpty'), trigger: 'change' },
    {
      validator: (rule, value, callback) => {
        if (value < 0 || value > 100) return callback(new Error('0-100'))
        if (form.value.maxPercentage !== null && value > form.value.maxPercentage) return callback(new Error('min<=max'))
        callback()
      },
      trigger: 'change'
    }
  ],
  maxPercentage: [
    { required: true, message: t('general.cannotBeEmpty'), trigger: 'change' },
    {
      validator: (rule, value, callback) => {
        if (value < 0 || value > 100) return callback(new Error('0-100'))
        if (form.value.minPercentage !== null && value < form.value.minPercentage) return callback(new Error('min<=max'))
        callback()
      },
      trigger: 'change'
    }
  ]
}

const tableData = ref([])
const total = ref(0)
const page = ref(1)
const pageSize = ref(10)
const multipleSelection = ref([])

const handleSizeChange = (val) => {
  pageSize.value = val
  fetchTable()
}

const handleCurrentChange = (val) => {
  page.value = val
  fetchTable()
}

const fetchTable = async () => {
  const res = await getRatingList({ page: page.value, pageSize: pageSize.value })
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

const resetForm = () => {
  form.value = {
    ID: undefined,
    rating: null,
    description: '',
    performanceLevel: '',
    minPercentage: 0,
    maxPercentage: 0
  }
}

const openDrawer = () => {
  type.value = 'create'
  resetForm()
  drawerVisible.value = true
}

const closeDrawer = () => {
  drawerVisible.value = false
  resetForm()
}

const onEdit = async (row) => {
  const res = await findRating({ ID: row.ID })
  if (res.code === 0) {
    const r = res.data.rerating
    form.value = {
      ID: r.ID,
      rating: r.rating,
      description: r.description,
      performanceLevel: r.performanceLevel,
      minPercentage: r.minPercentage,
      maxPercentage: r.maxPercentage
    }
    type.value = 'update'
    drawerVisible.value = true
  }
}

const submitForm = () => {
  formRef.value.validate(async (valid) => {
    if (!valid) return
    let res
    if (type.value === 'create') {
      res = await createRating(form.value)
    } else {
      res = await updateRating(form.value)
    }
    if (res.code === 0) {
      ElMessage.success(t('general.createUpdateSuccess'))
      closeDrawer()
      fetchTable()
    }
  })
}

const onDelete = (row) => async () => {
  const res = await deleteRating({ ID: row.ID })
  if (res.code === 0) {
    ElMessage.success(t('general.deleteSuccess'))
    if (tableData.value.length === 1 && page.value > 1) page.value--
    fetchTable()
  }
}

const onSelectionChange = (rows) => {
  multipleSelection.value = rows
}

const onBatchDelete = async () => {
  if (!multipleSelection.value.length) {
    ElMessage.warning(t('general.selectDataToDelete'))
    return
  }
  const ids = multipleSelection.value.map((i) => i.ID)
  const res = await deleteRatingByIds({ ids })
  if (res.code === 0) {
    ElMessage.success(t('general.deleteSuccess'))
    if (tableData.value.length === ids.length && page.value > 1) page.value--
    fetchTable()
  }
}

const renderRange = (min, max) => {
  if (min === 0) return `≤ ${max}%`
  return `${min}% - ${max}%`
}
</script>

<style scoped>
.gva-table-box { padding: 20px; }
.gva-btn-list { margin-bottom: 10px; }
</style>

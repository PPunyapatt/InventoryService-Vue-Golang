<template>
  <div class="container">

    <!-- รหัสสินค้า -->
    <v-row align="center">
      <v-col cols="2" style="margin-bottom: 20px;">
        <h3>รหัสสินค้า</h3>
      </v-col>
      <v-col cols="9">
        <v-text-field  
          label=""
          :rules="rules"
          v-model="state.text" 
          variant="outlined" 
          density="compact" 
          placeholder="XXXX-XXXX-XXXX-XXXX"
          @input="formatCode"
        />
      </v-col>
      <v-col cols="1" style="margin-bottom: 20px;">
        <v-btn 
            color="primary" 
            :disabled="state.text.replace(/-/g, '').length != 16"
            block
            @click="addInventoryCode"
        >ADD</v-btn>
      </v-col>
    </v-row>
    <v-table 
        class="full-height-table"
        ref="tableWrapper" 
        fixed-header 
        style="margin-top: 10px;">
      <thead>
        <tr>
          <th class="text-center">id</th>
          <th class="text-center" style="width: 35%">รหัสสินค้า (16 หลัก)</th>
          <th class="text-center">บาร์โค้ดสินค้า</th>
          <th class="text-center" style="width: 20%">Action</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(item, idx) in state.inventories" :key="item.name">
          <td class="text-center">{{ idx + 1 + (state.pagination.page-1)*state.pagination.limit }}</td>
          <td class="text-center">{{ formatName(item.inventory_code) }}</td>
          <td class="text-center">
            <Barcode :value="item.inventory_code" />
          </td>
          <td class="text-center">
            <v-btn 
              color="red" 
              size="small"
            @click="confirmDelete(item.inventory_code)">ลบ</v-btn>
          </td>

          
        </tr>
      </tbody>
    </v-table>

    <v-pagination
      v-model="state.pagination.page"
      :length="state.pagination.pageTotal"
      rounded="circle"
      :total-visible="5"
      @update:model-value="pageChange"
    >
      <template #prev>
        <v-btn icon @click="chevronChange('left')" :disabled="state.pagination.page === 1" rounded="circle">
          <v-icon>
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 640 640">
              <path d="M169.4 297.4C156.9 309.9 156.9 330.2 169.4 342.7L361.4 534.7C373.9 547.2 394.2 547.2 406.7 534.7C419.2 522.2 419.2 501.9 406.7 489.4L237.3 320L406.6 150.6C419.1 138.1 419.1 117.8 406.6 105.3C394.1 92.8 373.8 92.8 361.3 105.3L169.3 297.3z"/>
            </svg>
          </v-icon>
        </v-btn>
      </template>

      <template #next>
        <v-btn icon @click="chevronChange('right')" :disabled="state.pagination.page === state.pagination.pageTotal" rounded="circle">
          <v-icon>
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 640 640">
              <path d="M471.1 297.4C483.6 309.9 483.6 330.2 471.1 342.7L279.1 534.7C266.6 547.2 246.3 547.2 233.8 534.7C221.3 522.2 221.3 501.9 233.8 489.4L403.2 320L233.9 150.6C221.4 138.1 221.4 117.8 233.9 105.3C246.4 92.8 266.7 92.8 279.2 105.3L471.2 297.3z"/>
            </svg>
          </v-icon>
        </v-btn>
      </template>
    </v-pagination>

    <Confirm 
      v-model:open="state.open_confirm" 
      :text="`ต้องการลบข้อมูลรหัสสินค้า ${formatName(state.selectCode)} หรือไม่ ?`" 
      @confirm="deleteInventoryCode"
    />
     <Alert
      v-model:show="state.alert"
      title="Error"
      :message="state.message"
    />

    <Loading
      v-model:show="state.loading"
    />

    <SnackBar
      v-model:show="state.snackAlert" 
      :message="state.message"
    />
  </div>
</template>

<script setup>
import { reactive, onMounted, ref , getCurrentInstance } from 'vue'
import Barcode from '../components/Barcode.vue'
import InventoryService from "../services/api-services"
import Confirm from "../components/confirm.vue"
import Loading from "../components/Loading.vue"
import SnackBar from "../components/SnackBar.vue"

const state = reactive({ 
      text: "",
      loading: false,
      disable_add: false,
      open_confirm: false,
      inventories_more: [],
      inventories: [],
      pagination: {
        page: 1,
        limit: 10,
        total: 0,
        pageTotal: 0,
      },
      showSnackbarAlert: false,
      message: '',
      selectCode: "",
      alert: true,
      snackAlert: false
    })
  
    const tableWrapper = ref(null)

    const rules = [
      value => {
        if (!value) return true
        const cleaned = value.replace(/-/g, '')
        if (cleaned.length === 16) {
            state.disable_add = false
        }
        return cleaned.length === 16 || 'Must be 16 characters'
      }
    ]

    function formatCode() {
      state.text = state.text.toUpperCase()
      let cleaned = state.text.replace(/\W/g, '')
      if (cleaned.length > 16) cleaned = cleaned.slice(0, 16)
      const parts = cleaned.match(/.{1,4}/g)
      state.text = parts ? parts.join('-') : ''
    }

    function formatName(name) {
        if (name === '') {
          return name
        }
        const parts = name.match(/.{1,4}/g)
        return parts.join('-')
    }

    async function loadData() {
      state.loading = true
        try {
          var response = await InventoryService.getInventories(state.pagination.page, state.pagination.limit)
          var data = response.data
          state.inventories = data.data;
          state.pagination.total = data._pagination.total_count
          state.pagination.pageTotal =  Math.ceil(state.pagination.total / state.pagination.limit);
          
        } catch (e) {
          if (e?.response?.data?.error_message != undefined) {
              state.message = e.response.data.message
          } else {
              state.message = e.message
          }
          state.alert = true
        } finally {
          state.loading = false
        }
    }

    const deleteInventoryCode = async (val) => {
      console.log("val: ", val);
      
      if (val === 'cancel') return
        
      state.loading = true

      var data = {
        inventory_code : state.selectCode
      }

      try {
          const response = await InventoryService.deleteInventory(data)
          state.message = response.data.message
          state.pagination.total--
          state.pagination.pageTotal =  Math.ceil(state.pagination.total / state.pagination.limit);

          state.inventories = state.inventories.filter(x => x.inventory_code !== state.selectCode);

          if (state.pagination.page > state.pagination.pageTotal) {
            state.pagination.page = state.pagination.pageTotal || 1;
          } 
          
          if (state.inventories.length < state.pagination.limit) {
            await loadData()
          }

          state.snackAlert = true
      } catch (e) {
          if (e?.response?.data?.error_message != undefined) {
              state.message = e.response.data.message
          } else {
              state.message = e.message
          }
          state.alert = true
      } finally {
        state.loading = false
      }
    }

    const addInventoryCode = async () => {
      state.loading = true

      var data = {
        inventory_code : state.text.replace(/-/g, '')
      }

      try {
          const response = await InventoryService.addInventory(data)
          state.message = response.data.message
          state.pagination.total++
          state.pagination.pageTotal =  Math.ceil(state.pagination.total / state.pagination.limit);
          state.snackAlert = true

          if (state.inventories.length < state.pagination.limit &&  state.pagination.pageTotal === state.pagination.page) {
              await loadData()
          }
      } catch (e) {
          if (e?.response?.data?.error_message != undefined) {
              state.message = e.response.data.message
          } else {
              state.message = e.message
          }
          state.alert = true
      } finally {
        state.text = ""
        state.loading = false
      }
      
       
      
    }

    const confirmDelete = (inventoryCode) => {
      state.selectCode =  inventoryCode
      state.open_confirm = true
    }


    const pageChange = async (p) => {
      console.log("Page changed:", p);
      state.pagination.page = p
      await loadData()
    }

    const chevronChange = async (val) => {
      if (val === "left" && state.pagination.page > 1) {
        state.pagination.page--
      } else if (val === "right" && state.pagination.page < state.pagination.pageTotal) {
        state.pagination.page++
      }
      await loadData()
    }

    onMounted(async () => {
        await loadData(state.pagination.limit)
    })
</script>

<style>
.container {
  align-items: center;
  justify-content: center;
  min-height: 100vh;
  padding: 20px;
  max-width: 900px;
  width: 100%;
  box-sizing: border-box;
  margin: 30px auto;
}

th:nth-child(1) { width: 10%; }
th:nth-child(2) { width: 40%; }
th:nth-child(3) { width: 40%; }
th:nth-child(4) { width: 10%; }

.v-table.v-table--fixed-header > .v-table__wrapper > table > thead > tr > th {
  background-color: rgb(81, 199, 236) !important;
  font-size: 16px;
  font-weight: bold;
}

  .full-height-table {
    height: calc(100vh - 150px); /* ลบ header, input row หรือ margin-top อื่น ๆ */
  }
  .v-table.v-table--fixed-header > .v-table__wrapper {
    height: 100%;
  }

</style>

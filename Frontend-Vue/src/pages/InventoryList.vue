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
            :disabled="state.disable_add"
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

        <tr v-if="state.loading">
          <td colspan="4" class="text-center">
            <v-progress-circular color="primary" indeterminate></v-progress-circular>
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
        <v-btn icon @click="chevronChange('right')" :disabled="state.pagination.page === 20" rounded="circle">
          <v-icon>
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 640 640">
              <path d="M471.1 297.4C483.6 309.9 483.6 330.2 471.1 342.7L279.1 534.7C266.6 547.2 246.3 547.2 233.8 534.7C221.3 522.2 221.3 501.9 233.8 489.4L403.2 320L233.9 150.6C221.4 138.1 221.4 117.8 233.9 105.3C246.4 92.8 266.7 92.8 279.2 105.3L471.2 297.3z"/>
            </svg>
          </v-icon>
        </v-btn>
      </template>
    </v-pagination>

    <v-snackbar
      v-model="state.showSnackbarAlert"
      :timeout="3000"
      location="bottom  end"
      :color="state.color"
      elevation="2"
      :multi-line="false"
    >
      {{ state.message }}
    </v-snackbar>

    <Confirm 
      :open="state.open_confirm" 
      :text="`ต้องการลบข้อมูลรหัสสินค้า ${formatName(state.selectCode)} หรือไม่ ?`" 
      @confirm="deleteInventoryCode"
    />
     <v-overlay
      :model-value="state.loading"
      class="align-center justify-center"
    >
      <v-progress-circular
        color="primary"
        size="64"
        indeterminate
      ></v-progress-circular>
    </v-overlay>
  </div>
</template>

<script>
import { reactive, onMounted, ref , getCurrentInstance } from 'vue'
import Barcode from '../components/Barcode.vue'
import InventoryService from "../services/api-services"
import Confirm from "../components/confirm.vue"

export default {
  components: { 
    Barcode,
    Confirm
   },
  setup() {

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
      selectCode: ""
    })
    const rowHeight = 56; // ความสูงประมาณของแต่ละ row
    const tableHeight = window.innerHeight - 150; // ตามที่คุณตั้งค่า
    const rowsPerPage = Math.ceil(tableHeight / rowHeight);

    const tableWrapper = ref(null)

    const formatCode = () => {
      state.text = state.text.toUpperCase()
      let cleaned = state.text.replace(/\W/g, '')
      if (cleaned.length > 16) cleaned = cleaned.slice(0, 16)
      const parts = cleaned.match(/.{1,4}/g)
      state.text = parts ? parts.join('-') : ''
    }

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

    const loadData = () => {
      state.loading = true
      InventoryService.getInventories(state.pagination.page, state.pagination.limit)
        .then(response => {
          state.pagination.total = response.data._pagination.total_count
          state.inventories = response.data.data;
          state.pagination.pageTotal =  Math.ceil(state.pagination.total / state.pagination.limit);
        })
        .catch(e => {
          console.log(e);
        });
      state.loading = false
    }

    const formatName = (name) => {
        if (name === '') {
          return name
        }
        const parts = name.match(/.{1,4}/g)
        return parts.join('-')
    }

    const deleteInventoryCode = async (val) => {
      if (val === 'cancel') return
      state.open_confirm = false
      state.loading = true

      var data = {
        inventory_code : state.selectCode
      }

      try {
          const response = await InventoryService.deleteInventory(data)
          state.message = "Delete Success"
          state.color = 'success'
          state.pagination.total--
          state.pagination.pageTotal =  Math.ceil(state.pagination.total / state.pagination.limit);

          state.inventories = state.inventories.filter(x => x.inventory_code !== state.selectCode);

          if (state.pagination.page > state.pagination.pageTotal) {
            state.pagination.page = state.pagination.pageTotal || 1;
          } 
          
          if (state.inventories.length < state.pagination.limit) {
            await loadData()
          }
      } catch (e) {
          state.color = 'red'
          state.message = e.response.data.error_message
      } finally {
        state.showSnackbarAlert = true
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
          state.message = "Add inventory code success"
          state.color = 'success'
          state.pagination.total++
          state.pagination.pageTotal =  Math.ceil(state.pagination.total / state.pagination.limit);
      } catch (e) {
          state.color = 'red'
          state.message = e.response.data.error_message
      } finally {
        state.showSnackbarAlert = true
        state.text = ""
        state.disable_add = true
        state.loading = false
      }
      
       
      if (state.inventories.length > state.pagination.limit &&  state.pagination.pageTotal === state.pagination.page) {
        await loadData()
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
      } else if (val === "right" && state.pagination.page < state.pagination.total) {
        state.pagination.page++
      }
      console.log("page: ", state.pagination.page);
      await loadData()
    }



    onMounted(async () => {
        await loadData(state.pagination.limit)
      })

    return { 
      state,
      rules, 
      tableWrapper,
      formatCode, 
      formatName,
      loadData,
      deleteInventoryCode,
      addInventoryCode,
      confirmDelete,
      pageChange,
      chevronChange
    }
  }
}
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

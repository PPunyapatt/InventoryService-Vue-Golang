import http from "../http"

class InventoryService {
    getInventories(page, limit) {
        return http.get(`/inventories/?page=${page}&limit=${limit}`)
    }

    addInventory(data) {
        return http.post("/inventory",data);
    }

    deleteInventory(data) {
        return http.delete(`/inventory`, { data: data })
    }
}

export default new InventoryService();
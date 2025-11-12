import http from "../http"

class InventoryService {
    getInventories(page, limit) {
        return http.get(`/inventories/?page=${page}&limit=${limit}`)
    }

    addInventory(data) {
        return http.post("/tutorials", data);
    }

    deleteInventory(inventory_code) {
        return http.delete(`inventory/${inventory_code}`)
    }
}

export default new InventoryService();
import { EventListResponse, EventResponse } from "@/types/event"
import { Api } from "./api"

async function createOne(name: string, location: string, date: string): Promise<EventResponse> {
    return Api.post("/event", { name, location, date })
}

async function getOne(id: string): Promise<EventResponse> {
    return Api.get(`/event/${id}`)
}

async function getAll(): Promise<EventListResponse> {
    return Api.get("/event")
}

async function updateOne(id: string, name: string, location: string, date: string): Promise<EventResponse> {
    return Api.put(`/event/${id}`, { name, location, date })
}

async function deleteOne(id: string): Promise<null> {
    return Api.delete(`/event/${id}`)
}

const eventService = {
    createOne,
    getOne,
    getAll,
    updateOne,
    deleteOne,
}

export { eventService }
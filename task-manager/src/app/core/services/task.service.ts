import { HttpClient } from "@angular/common/http";
import { environment } from "../../../environments/environment";
import { Observable } from "rxjs";
import { CreateTaskPayload, GetTasksResponse, Task } from "../models/task.interface";
import { EventEmitter, Injectable, Output } from "@angular/core";

@Injectable({
  providedIn: "root",
})
export class TaskService {
  private readonly baseUrl = environment.apiBaseUrl + '/tasks/'

  constructor(private http: HttpClient) {}

  getTasks(): Observable<GetTasksResponse> {
    return this.http.get<GetTasksResponse>(this.baseUrl)
  }

  addTask(task: CreateTaskPayload): Observable<Task> {
    return this.http.post<Task>(this.baseUrl ,task)
  }

  @Output() public dataStatus: EventEmitter<any> = new EventEmitter()
}
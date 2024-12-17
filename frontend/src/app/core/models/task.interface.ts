export interface Task {
  ID:          string;
  Title:       string;
  Description: string;
  CreatedAt:   string;
}

export interface CreateTaskPayload {
  title:       string;
  description: string;
}

export type GetTasksResponse = Task[]
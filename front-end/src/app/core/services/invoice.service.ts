import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { CreateInvoiceRequest, Invoice } from "../models/invoice.model";

@Injectable({ providedIn: 'root' })
export class InvoiceService {
  private readonly url = '/api/invoices';

  constructor(private http: HttpClient) {}

  getAll(): Observable<Invoice[]> {
    return this.http.get<Invoice[]>(this.url);
  }

  create(payload: CreateInvoiceRequest): Observable<Invoice> {
    return this.http.post<Invoice>(this.url, payload);
  }

  // RxJS: usado com switchMap na tela de lista para recarregar após imprimir
  print(id: number): Observable<Invoice> {
    return this.http.post<Invoice>(`${this.url}/${id}/print`, {});
  }
}
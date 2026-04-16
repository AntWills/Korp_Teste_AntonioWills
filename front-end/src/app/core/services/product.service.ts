import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";
import { CreateProductRequest, Product } from "../models/product.model";

@Injectable({ providedIn: 'root' })
export class ProductService {
  private readonly url = '/api/inventory/products';

  constructor(private http: HttpClient) {}

  // RxJS: Observable retornado diretamente — o componente faz o subscribe
  getAll(): Observable<Product[]> {
    return this.http.get<Product[]>(this.url);
  }

  create(payload: CreateProductRequest): Observable<Product> {
    return this.http.post<Product>(this.url, payload);
  }
}
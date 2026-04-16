import { Component, OnDestroy, OnInit } from "@angular/core";
import { Product } from "../../../core/models/product.model";
import { Subject, takeUntil } from "rxjs";
import { ProductService } from "../../../core/services/product.service";
import { MatSnackBar } from "@angular/material/snack-bar";

@Component({ selector: 'app-product-list', templateUrl: './product-list.component.html' })
export class ProductListComponent implements OnInit, OnDestroy {
  products: Product[] = [];
  isLoading = false;
  displayedColumns = ['code', 'description', 'balance'];

  // RxJS: Subject usado com takeUntil para cancelar subscriptions no destroy
  private destroy$ = new Subject<void>();

  constructor(
    private productService: ProductService,
    private snackBar: MatSnackBar
  ) {}

  // Ciclo de vida: ngOnInit — carrega dados ao inicializar o componente
  ngOnInit(): void {
    this.loadProducts();
  }

  loadProducts(): void {
    this.isLoading = true;
    this.productService.getAll()
      .pipe(takeUntil(this.destroy$))   // RxJS: cancela se componente for destruído
      .subscribe({
        next: (data) => { this.products = data; this.isLoading = false; },
        error: () => { this.isLoading = false; }
        // erro global tratado pelo ErrorInterceptor
      });
  }

  // Ciclo de vida: ngOnDestroy — evita memory leak
  ngOnDestroy(): void {
    this.destroy$.next();
    this.destroy$.complete();
  }
}
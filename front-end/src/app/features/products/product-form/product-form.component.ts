import { Component, EventEmitter, Output } from "@angular/core";
import { FormBuilder, Validators } from "@angular/forms";
import { CreateProductRequest, Product } from "../../../core/models/product.model";
import { ProductService } from "../../../core/services/product.service";
import { MatSnackBar } from "@angular/material/snack-bar";

@Component({ selector: 'app-product-form', templateUrl: './product-form.component.html' })
export class ProductFormComponent {
  isSubmitting = false;

  // Reactive Forms com validação
  form = this.fb.group({
    code:        ['', [Validators.required, Validators.maxLength(20)]],
    description: ['', [Validators.required, Validators.maxLength(80)]],
    balance:     [0,  [Validators.required, Validators.min(0)]]
  });

  // RxJS: EventEmitter para notificar o componente pai de novo produto criado
  @Output() productCreated = new EventEmitter<Product>();

  constructor(
    private fb: FormBuilder,
    private productService: ProductService,
    private snackBar: MatSnackBar
  ) {}

  submit(): void {
    if (this.form.invalid) { this.form.markAllAsTouched(); return; }
    this.isSubmitting = true;

    this.productService.create(this.form.value as CreateProductRequest)
      .subscribe({
        next: (product) => {
          this.snackBar.open('Produto cadastrado!', 'Fechar', { duration: 3000 });
          this.form.reset({ balance: 0 });
          this.isSubmitting = false;
          this.productCreated.emit(product);  // notifica o pai para recarregar a tabela
        },
        error: () => { this.isSubmitting = false; }
      });
  }
}
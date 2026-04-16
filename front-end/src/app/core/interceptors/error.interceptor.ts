import { HttpErrorResponse, HttpEvent, HttpHandler, HttpInterceptor, HttpRequest } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { MatSnackBar } from "@angular/material/snack-bar";
import { catchError, Observable, throwError } from "rxjs";

// core/interceptors/error.interceptor.ts
@Injectable()
export class ErrorInterceptor implements HttpInterceptor {
  constructor(private snackBar: MatSnackBar) {}

  intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    return next.handle(req).pipe(
      // RxJS: catchError para capturar erros HTTP globalmente
      catchError((error: HttpErrorResponse) => {
        let message = 'Erro inesperado';

        if (error.status === 0) {
          message = 'Serviço indisponível. Verifique sua conexão.';
        } else if (error.status === 409) {
          message = error.error?.error ?? 'Conflito ao processar a operação.';
        } else if (error.status === 503) {
          message = 'Serviço de estoque indisponível no momento.';
        } else if (error.error?.error) {
          message = error.error.error;
        }

        this.snackBar.open(message, 'Fechar', { duration: 5000 });
        return throwError(() => error);
      })
    );
  }
}
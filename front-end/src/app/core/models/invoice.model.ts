export type InvoiceStatus = 'Aberta' | 'Fechada';

export interface InvoiceItem {
  productId: number;
  quantity: number;
}

export interface Invoice {
  id: number;
  number: number;
  status: InvoiceStatus;
  items: InvoiceItem[];
  createdAt: string;
}

export interface CreateInvoiceRequest {
  items: InvoiceItem[];
}
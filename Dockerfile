# ใช้ official Golang image
FROM golang:1.23.3-alpine AS builder

# ติดตั้ง dependencies ที่จำเป็น
RUN apk add --no-cache gcc musl-dev

# ตั้งค่า working directory
WORKDIR /app

# คัดลอกไฟล์ทั้งหมดเข้า container
COPY . .

# ดาวน์โหลด dependencies และ build
RUN go mod tidy
RUN go build -o main .

# ใช้ image ที่เล็กลงสำหรับ production
FROM alpine:latest

# ตั้งค่า working directory
WORKDIR /root/

# คัดลอกไฟล์จาก builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/.env . 

# เปิด port (เช่น 8080)
EXPOSE 8080

# รัน main.go
CMD ["./main"]

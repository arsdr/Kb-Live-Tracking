document.addEventListener('DOMContentLoaded', (event) => {
          const transactionsBody = document.getElementById('transactions-body');
          const headerRow = transactionsBody.querySelector('tr'); // Get the header row
          const ws = new WebSocket('ws://localhost/ws/transactions');

          ws.onmessage = function (event) {
                    const newTransaction = JSON.parse(event.data);
                    addTransactionRow(newTransaction);
          };

          function addTransactionRow(transaction) {
                    const row = document.createElement('tr');

                    const timestampCell = document.createElement('td');
                    timestampCell.textContent = transaction.date;
                    row.appendChild(timestampCell);

                    const idCell = document.createElement('td');
                    idCell.textContent = transaction.id_trx;
                    row.appendChild(idCell);

                    const whatsappCell = document.createElement('td');
                    whatsappCell.textContent = transaction.contact_no;
                    row.appendChild(whatsappCell);

                    const amountCell = document.createElement('td');
                    amountCell.textContent = transaction.price;
                    row.appendChild(amountCell);

                    const statusCell = document.createElement('td');
                    const statusBadge = document.createElement('span');
                    statusBadge.textContent = transaction.status;
                    statusBadge.classList.add('badge', transaction.status === 'Success' ? 'bg-success' : 'bg-danger');
                    statusCell.appendChild(statusBadge);
                    row.appendChild(statusCell);


                    transactionsBody.insertBefore(row, headerRow.nextSibling);

                    while (transactionsBody.querySelectorAll('tr').length > 11) {
                              transactionsBody.removeChild(transactionsBody.lastChild);
                    }
          }

          ws.onclose = function () {
                    console.log('WebSocket connection closed');
          };
});

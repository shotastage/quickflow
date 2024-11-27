// File: internal/interfaces/httpserver/handler/status_handler.go

package handler

import (
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
)

type StatusHandler struct {
	templates *template.Template
}

func NewStatusHandler() *StatusHandler {
	return &StatusHandler{
		templates: template.Must(template.New("status").Parse(statusPageTemplate)),
	}
}

func (h *StatusHandler) HandleStatusPage(c echo.Context) error {
	return c.HTML(http.StatusOK, statusPageTemplate)
}

const statusPageTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Quickflow Zone 1</title>
    <link href="https://cdnjs.cloudflare.com/ajax/libs/tailwindcss/2.2.19/tailwind.min.css" rel="stylesheet">
</head>
<body class="bg-gray-50">
    <div class="min-h-screen">
        <!-- Top Navigation -->
        <nav class="bg-white shadow-sm">
            <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
                <div class="flex justify-between h-14">
                    <div class="flex items-center">
                        <span class="text-lg font-semibold text-gray-900">zone1.quickflow.com</span>
                    </div>
                    <div class="flex items-center space-x-4">
                        <a href="https://status.quickflow.com" class="text-sm text-gray-500 hover:text-gray-900">Status Page</a>
                        <a href="/login" class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700">
                            Login
                        </a>
                    </div>
                </div>
            </div>
        </nav>

        <!-- Main Content -->
        <div class="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
            <!-- System Status Overview -->
            <div class="bg-white shadow rounded-lg mb-6">
                <div class="px-4 py-5 sm:p-6">
                    <div class="flex items-center" id="systemStatus">
                        <div class="flex-shrink-0">
                            <div class="rounded-full h-4 w-4 bg-gray-300" id="statusIndicator"></div>
                        </div>
                        <div class="ml-3">
                            <h3 class="text-lg font-medium text-gray-900" id="statusMessage">Checking System Status...</h3>
                            <p class="text-sm text-gray-500" id="lastChecked">Checking...</p>
                        </div>
                    </div>
                </div>
            </div>

            <!-- Service Status -->
            <div class="bg-white shadow rounded-lg overflow-hidden mb-6">
                <div class="px-4 py-5 sm:px-6">
                    <h3 class="text-lg font-medium text-gray-900">Service Status</h3>
                </div>
                <div class="border-t border-gray-200">
                    <dl>
                        <div class="bg-gray-50 px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
                            <dt class="text-sm font-medium text-gray-500">API Server</dt>
                            <dd class="mt-1 text-sm text-gray-900 sm:mt-0 sm:col-span-2 flex items-center">
                                <span class="rounded-full h-3 w-3 bg-gray-300 mr-2" id="serverHealthIndicator"></span>
                                <span id="serverHealthStatus">Checking...</span>
                            </dd>
                        </div>
                        <div class="bg-white px-4 py-5 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
                            <dt class="text-sm font-medium text-gray-500">Database</dt>
                            <dd class="mt-1 text-sm text-gray-900 sm:mt-0 sm:col-span-2 flex items-center">
                                <span class="rounded-full h-3 w-3 bg-gray-300 mr-2" id="dbHealthIndicator"></span>
                                <span id="dbHealthStatus">Checking...</span>
                            </dd>
                        </div>
                    </dl>
                </div>
            </div>

            <!-- Incident History (Static content) -->
            <div class="bg-white shadow rounded-lg overflow-hidden">
                <div class="px-4 py-5 sm:px-6">
                    <h3 class="text-lg font-medium text-gray-900">Recent Incidents</h3>
                </div>
                <div class="border-t border-gray-200">
                    <div class="flow-root">
                        <ul class="divide-y divide-gray-200">
                            <li class="px-4 py-5">
                                <div class="flex items-center space-x-4">
                                    <div class="flex-shrink-0">
                                        <span class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full bg-yellow-100 text-yellow-800">Resolved</span>
                                    </div>
                                    <div class="flex-1 min-w-0">
                                        <p class="text-sm font-medium text-gray-900">Increased API Latency</p>
                                        <p class="text-sm text-gray-500">Nov 24, 2024 - Resolved within 15 minutes</p>
                                    </div>
                                </div>
                                <div class="mt-2">
                                    <p class="text-sm text-gray-500">Brief period of increased latency due to scheduled database maintenance. All systems operating normally now.</p>
                                </div>
                            </li>
                            <li class="px-4 py-5">
                                <div class="flex items-center space-x-4">
                                    <div class="flex-shrink-0">
                                        <span class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full bg-green-100 text-green-800">Resolved</span>
                                    </div>
                                    <div class="flex-1 min-w-0">
                                        <p class="text-sm font-medium text-gray-900">Planned Maintenance</p>
                                        <p class="text-sm text-gray-500">Nov 20, 2024 - Completed as scheduled</p>
                                    </div>
                                </div>
                                <div class="mt-2">
                                    <p class="text-sm text-gray-500">Scheduled system upgrade completed successfully. No service interruption occurred.</p>
                                </div>
                            </li>
                        </ul>
                    </div>
                </div>
            </div>
        </div>

        <!-- Footer -->
        <footer class="bg-white mt-8 border-t border-gray-200">
            <div class="max-w-7xl mx-auto py-6 px-4 sm:px-6 lg:px-8">
                <div class="text-center text-sm text-gray-500">
                    <p>For immediate support, please contact: support@quickflow.com</p>
                    <p class="mt-1">Subscribe to real-time update for system: <a href="/health" class="text-indigo-600 hover:text-indigo-500">Status Updates</a></p>
                </div>
            </div>
        </footer>
    </div>

    <script>
        async function checkSystemStatus() {
            try {
                const response = await fetch('/health');
                const data = await response.json();

                const statusIndicator = document.getElementById('statusIndicator');
                const statusMessage = document.getElementById('statusMessage');
                const lastChecked = document.getElementById('lastChecked');
                const serverHealthIndicator = document.getElementById('serverHealthIndicator');
                const serverHealthStatus = document.getElementById('serverHealthStatus');
                const dbHealthIndicator = document.getElementById('dbHealthIndicator');
                const dbHealthStatus = document.getElementById('dbHealthStatus');

                // Update server status
                if (data.server.status === 'ok') {
                    serverHealthIndicator.className = 'rounded-full h-3 w-3 bg-green-500 mr-2';
                    serverHealthStatus.textContent = data.server.message || 'Operational';
                } else {
                    serverHealthIndicator.className = 'rounded-full h-3 w-3 bg-red-500 mr-2';
                    serverHealthStatus.textContent = data.server.message || 'Issues Detected';
                }

                // Update database status
                if (data.database.status === 'ok') {
                    dbHealthIndicator.className = 'rounded-full h-3 w-3 bg-green-500 mr-2';
                    dbHealthStatus.textContent = data.database.message || 'Operational';
                } else {
                    dbHealthIndicator.className = 'rounded-full h-3 w-3 bg-red-500 mr-2';
                    dbHealthStatus.textContent = data.database.message || 'Issues Detected';
                }

                // Update overall system status
                const isAllOk = data.server.status === 'ok' && data.database.status === 'ok';
                if (isAllOk) {
                    statusIndicator.className = 'rounded-full h-4 w-4 bg-green-500';
                    statusMessage.textContent = 'All Systems Operational';
                } else {
                    statusIndicator.className = 'rounded-full h-4 w-4 bg-red-500';
                    statusMessage.textContent = 'System Issues Detected';
                }

                // Update last checked time
                const now = new Date();
                lastChecked.textContent = 'Last checked: ' + now.toLocaleTimeString();

            } catch (error) {
                console.error('Error checking system status:', error);
                // Update UI to show error state
                document.getElementById('statusIndicator').className = 'rounded-full h-4 w-4 bg-red-500';
                document.getElementById('statusMessage').textContent = 'Unable to Check System Status';
                document.getElementById('serverHealthIndicator').className = 'rounded-full h-3 w-3 bg-red-500 mr-2';
                document.getElementById('serverHealthStatus').textContent = 'Unable to Connect';
                document.getElementById('dbHealthIndicator').className = 'rounded-full h-3 w-3 bg-red-500 mr-2';
                document.getElementById('dbHealthStatus').textContent = 'Unable to Connect';
            }
        }

        // Check status immediately and then every 30 seconds
        checkSystemStatus();
        setInterval(checkSystemStatus, 30000);
    </script>
</body>
</html>`
